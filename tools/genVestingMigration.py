from os.path import join
import os
import errno

template = """case "{0}":
		sc, err := {0}.Connect(addressFile, cli)
		if err != nil {{
			return err
		}}

		_, err = cli.Transact(func(opts *bind.TransactOpts) (*types.Transaction, error) {{
			return sc.AddBeneficiaries(opts, addresses, amounts)
		}})
		return err
"""

outputDir = "../backend/contract/service/"
pools = [
    "Advisory",
    "IDO",
    "Liquidity",
    "Marketing",
    "Play2Earn",
    "Private",
    "Reserve",
    "Staking",
    "StrategicPartner",
    "Team"
]


for pool in pools:
    print(template.format(pool.lower()))
