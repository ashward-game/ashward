package openbox

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"orbit_nft/cmd/common"
	"orbit_nft/constant"
	"orbit_nft/contract"
	"orbit_nft/db/repository"
	"orbit_nft/db/service"
	"orbit_nft/nft/metadata"
	"orbit_nft/openbox"
	"orbit_nft/storage"
	"orbit_nft/storage/ipfs"
	"orbit_nft/storage/localstorage"
	"orbit_nft/util"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	cm "github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type OpenboxService struct {
	opener  *openbox.BoxOpener
	service *service.MakeRandService
	baseDir string
}

type HashWithClientRandom struct {
	Hash         cm.Hash
	ClientRandom []byte
	BoxGrade     int
}

// ServiceCmd represents the service command
var ServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Run openboxgenesis service as an RPC server",
	Long:  `Run openboxgenesis service as an RPC server.`,
	Run: func(cmd *cobra.Command, args []string) {
		baseDir, err := cmd.Flags().GetString("assets")
		if err != nil {
			log.Fatal(err)
		}
		chainId, err := cmd.Flags().GetString("chainId")
		if err != nil {
			log.Fatal(err)
		}
		secretFile, err := cmd.Flags().GetString("backendWallet")
		if err != nil {
			log.Fatal(err)
		}
		addressFile, err := cmd.Flags().GetString("addressFile")
		if err != nil {
			log.Fatal(err)
		}
		useIPFS, err := cmd.Flags().GetBool("useIPFS")
		if err != nil {
			log.Fatal(err)
		}
		mnemonic, err := util.GetSecrets(secretFile, "mnemonic")
		if err != nil {
			log.Fatal(err)
		}
		boxesConfigFile, err := cmd.Flags().GetString("boxesConfig")
		if err != nil {
			log.Fatal(err)
		}

		openboxServiceAddress := os.Getenv("OPENBOX_SERVICE_ADDRESS")
		var storage storage.Storage
		var baseURIToken string
		if useIPFS {
			ipfs, err := ipfs.NewShell(os.Getenv("IPFS_SHELL_URI"))
			if err != nil {
				log.Fatal(err)
			}
			baseURIToken = os.Getenv("IPFS_URI")
			storage = ipfs
		} else {
			local, err := localstorage.New(os.Getenv("LOCAL_STORAGE_DIR"))
			if err != nil {
				log.Fatal(err)
			}
			baseURIToken = os.Getenv("LOCAL_STORAGE_URI")
			storage = local
		}

		logger := common.NewLogger()
		ob, err := NewOpenBoxService(baseDir, baseURIToken, mnemonic, addressFile, chainId, storage, logger, boxesConfigFile)
		if err != nil {
			log.Fatal(err)
		}

		rpc.Register(ob)
		rpc.HandleHTTP()
		l, e := net.Listen("tcp", openboxServiceAddress)
		if e != nil {
			log.Fatal("listen error:", e)
		}
		go http.Serve(l, nil)
		log.Println("Openbox service is running ...")

		ch := make(chan os.Signal, 1)
		// kill (no param) default send syscall.SIGTERM
		// kill -2 is syscall.SIGINT
		// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		log.Println("Shutting down openbox service...")
	},
}

func init() {
	ServiceCmd.Flags().String("assets", "../assets", "Assets resource path for NFTs")
	ServiceCmd.Flags().String("chainId", "97", "Chain Id, default is testnet with chain id 97")
	ServiceCmd.Flags().String("backendWallet", "./secrets/wallet.json", "Secrets file for backend's account")
	ServiceCmd.Flags().String("addressFile", "../common/address.json", "Contracts' address file")
	ServiceCmd.Flags().Bool("useIPFS", false, "Use IPFS as storage")
	ServiceCmd.Flags().String("boxesConfig", "./config/openboxgenesis.json", "Boxes' configuration file")
}

func NewOpenBoxService(baseDir, baseURIToken, mnemonic, addressFile, chainId string, storage storage.Storage, logger *zap.Logger, boxesConfigFile string) (*OpenboxService, error) {
	sqlDB := common.MySQL()
	repo := repository.NewMakeRandRepository(sqlDB)
	service := service.NewMakeRandService(repo)
	secrets := contract.NewClientSecret(mnemonic, constant.EthDerivationPath)
	cli, err := contract.NewBscAuthenticatedClient(addressFile, chainId, secrets, contract.WithLogger(logger))
	if err != nil {
		return nil, err
	}

	metadt := metadata.NewMetadata(baseURIToken, storage)
	ob, err := openbox.NewBoxOpener(cli, metadt, boxesConfigFile)
	if err != nil {
		return nil, err
	}
	return &OpenboxService{
		opener:  ob,
		service: service,
		baseDir: baseDir,
	}, nil
}

func (b *OpenboxService) OpenBox(params *HashWithClientRandom, ignore *int) error {
	serverRandom, err := b.service.Reveal(params.Hash.Hex())
	if err != nil {
		return err
	}
	rarityPath := filepath.Join(b.baseDir, metadata.NftCharacter)
	return b.opener.OpenBox(rarityPath, params.BoxGrade, [32]byte(params.Hash), serverRandom, params.ClientRandom)
}
