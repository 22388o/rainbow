package main

import (
	"context"
	"fmt"

	bin "github.com/gagliardetto/binary"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	sol "github.com/streamingfast/solana-go"
)

const PsyOptionsProgramID = "R2y9ip6mxmWUj4pt54jP2hz2dgvMozy9VTSwMWE7evs"
const test = "BPFLoaderUpgradeab1e11111111111111111111111"
const USDCSolana = "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v"

func main() {
	pub := solana.MustPublicKeyFromBase58(PsyOptionsProgramID)
	endpoint := rpc.MainNetBeta_RPC
	client := rpc.New(endpoint)

	out, err := client.GetProgramAccounts(
		context.TODO(),
		pub,
	)
	if err != nil {
		panic(err)
	}
	//spew.Dump(len(out[0:3]))
	for _, i := range out[0:4] {
		spew.Dump(i.Pubkey)
		/*out, err := client.GetAccountInfo(
			context.TODO(),
			i.Pubkey,
		)
		spew.Dump(out)
		if err != nil {
			panic(err)
		}*/
		opt := new(OptionMarket)
		err = bin.NewBorshDecoder(i.Account.Data.GetBinary()).Decode(&opt)
		if err != nil {
			panic(err)
		}
		//spew.Dump(opt)
		a, b, c := deriveSerumMarketAddress(i.Pubkey, solana.PublicKey(opt.QuoteAssetMint), pub)
		spew.Dump(a, b, c)
		a, b, c = deriveSerumMarketAddress(i.Pubkey, solana.PublicKey(opt.UnderlyingAssetMint), pub)
		spew.Dump(a, b, c)
	}
	//f()
	//solana.MustPublicKeyFromBase58(test))
	/*res, err := client.GetAccountInfo(
		context.TODO(),
		pub,
	)
	if err != nil {
		panic(err)
	}
	//spew.Dump(len(res))
	spew.Dump(res)*/
	//psy_american.SetProgramID(pub)
	//fmt.Println(psy_american.ProgramName)
}
func deriveSerumMarketAddress(id, quote, programid solana.PublicKey) (solana.PublicKey, uint8, error) {
	seed := [][]byte{
		id[:],
		quote[:],
		[]byte("serumMarket"),
	}
	return solana.FindProgramAddress(seed, programid)
}

func dd(pub, quote, programid solana.PublicKey) (solana.PublicKey, uint8, error) {
	seed := [][]byte{
		[]byte(""),
		{1},
	}
	return solana.FindProgramAddress(seed, solana.MustPublicKeyFromBase58(test))
}

func f() {
	program_id := solana.MustPublicKeyFromBase58("BPFLoader1111111111111111111111111111111111")
	public_key := solana.MustPublicKeyFromBase58("SeedPubey1111111111111111111111111111111111")

	got, _ := solana.CreateProgramAddress([][]byte{
		{},
		{1},
	},
		program_id,
	)
	fmt.Println(got)

	got, _ = solana.CreateProgramAddress([][]byte{
		[]byte(""),
		{1},
	},
		program_id,
	)
	fmt.Println(got)

	got, _ = solana.CreateProgramAddress([][]byte{
		[]byte(""),
		//{1},
	},
		program_id,
	)
	fmt.Println(got)

	got, _ = solana.CreateProgramAddress([][]byte{
		[]byte("☉"),
		{1},
	},
		program_id,
	)
	fmt.Println(got)

	got, _ = solana.CreateProgramAddress([][]byte{
		[]byte("☉"),
	},
		program_id,
	)
	fmt.Println(got)

	got, _ = solana.CreateProgramAddress([][]byte{
		public_key[:],
		{1},
	},
		program_id,
	)
	fmt.Println(got)

	got, _ = solana.CreateProgramAddress([][]byte{
		[]byte("Talking"),
		[]byte("Squirrels"),
	},
		program_id,
	)
	fmt.Println(got)

}

/*var address solana.PublicKey
	var err error
	bumpSeed := uint8(math.MaxUint8)
	for bumpSeed > 0 {
		fmt.Print(bumpSeed)

		address, err = solana.CreateProgramAddress(append(seed, []byte{byte(bumpSeed)}), programID)
		if err == nil {
			fmt.Println(address)
		}
		bumpSeed--

	}
}*/

type anchorOptions struct {
	Options OptionMarket
}

type OptionMarket struct {
	OptionMint                  sol.PublicKey
	WriterTokenMint             sol.PublicKey
	UnderlyingAssetMint         sol.PublicKey
	QuoteAssetMint              sol.PublicKey
	UnderlyingAmountPerContract uint64
	QuoteAmountPerContract      uint64
	ExpirationUnixTimestamp     int64
	UnderlyingAssetPool         sol.PublicKey
	QuoteAssetPool              sol.PublicKey
	MintFeeAccount              sol.PublicKey
	ExerciseFeeAccount          sol.PublicKey
	Expired                     bool
	BumpSeed                    uint8
}
