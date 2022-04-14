package sign

import (
	"crypto/ecdsa"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"orbit_nft/cmd/common"
	"orbit_nft/crypto"
	"orbit_nft/db"
	"orbit_nft/db/repository"
	"orbit_nft/db/service"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"
)

// ServiceCmd represents the service command
var ServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Signing service via RPC",
	Long:  `Signing service via RPC.`,
	Run: func(cmd *cobra.Command, args []string) {
		keyPath := cmd.Flag("signingKey").Value.String()
		sqlDB := common.MySQL()
		signingServiceAddress := os.Getenv("SIGNING_SERVICE_ADDRESS")

		ss, err := NewSigningService(keyPath, sqlDB)
		if err != nil {
			log.Fatal(err)
		}
		rpc.Register(ss)
		rpc.HandleHTTP()
		l, e := net.Listen("tcp", signingServiceAddress)
		if e != nil {
			log.Fatal("listen error:", e)
		}
		go http.Serve(l, nil)
		log.Println("Signing service is running ...")

		ch := make(chan os.Signal, 1)
		// kill (no param) default send syscall.SIGTERM
		// kill -2 is syscall.SIGINT
		// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		log.Println("Shutting down signing service...")
	},
}

func init() {
	ServiceCmd.Flags().String("signingKey", "./secrets/sign.private", "Path to the signing key")
}

type SigningService struct {
	makeRandService *service.MakeRandService
	privateKey      *ecdsa.PrivateKey
}

func NewSigningService(keypath string, db *db.Database) (*SigningService, error) {
	sk, err := crypto.GetPrivateKey(keypath)
	if err != nil {
		return nil, err
	}

	makeRandRepo := repository.NewMakeRandRepository(db)
	makeRandService := service.NewMakeRandService(makeRandRepo)
	return &SigningService{
		makeRandService: makeRandService,
		privateKey:      sk,
	}, nil
}

func (ss *SigningService) Commit(ignored int, com *crypto.CommitmentWithSig) error {
	random, err := crypto.MakeRand()
	if err != nil {
		return err
	}

	sig, err := crypto.EthSign(ss.privateKey, random)
	if err != nil {
		return err
	}

	hash := crypto.HashMessage(random)
	err = ss.makeRandService.Commit(hexutil.Encode(hash), hexutil.Encode(random))
	if err != nil {
		return err
	}

	*com = crypto.CommitmentWithSig{
		Commitment: hexutil.Encode(hash),
		Signature:  hexutil.Encode(sig),
	}

	return nil
}
