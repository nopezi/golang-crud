package services

import (
	"eform-gateway/lib"
	"eform-gateway/requests"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

type Data struct {
	IsExistingCustomer    string
	JenisTransaksiNasabah string
	JenisForm             string
	NoRekeningTujuan      string
	NamaRekeningTujuan    string
	NamaPenyetor          string
	Keterangan            string
	TujuanPenyetoran      string
	JumlahSetor           string
	TermandCondition      string
	CustomerAgreement     string
	RequestBy             string
}

var ModuleTest = fx.Options(
	Module,
	fx.Invoke(Loop),
)

// func boot(
// 	lifecycle fx.Lifecycle,
// 	env lib.Env,
// 	elastic lib.Elasticsearch,
// ) {

// }
func Loop(repeated int) (repeatedCount int) {
	_ = godotenv.Load()

	url := os.Getenv("DBEHost")
	username := os.Getenv("DBEUsername")
	password := os.Getenv("DBEPassword")

	_, err := lib.New([]string{url}, username, password)
	if err != nil {
		fmt.Println(err)
	}
	transaction := TransactionService{}
	data := Data{
		IsExistingCustomer:    "Y",
		JenisTransaksiNasabah: "Setor Tunai",
		JenisForm:             "OPS02",
		NoRekeningTujuan:      "098917340986729",
		NamaRekeningTujuan:    "test",
		NamaPenyetor:          "test juga",
		Keterangan:            "test juga",
		TujuanPenyetoran:      "test juga",
		JumlahSetor:           "test juga",
		TermandCondition:      "test juga",
		CustomerAgreement:     "test juga",
		RequestBy:             "test juga",
	}

	for i := 0; i < repeated; i++ {
		_, _ = transaction.CreateTransaction(requests.TransactionRequest{
			Appname:     "BRIQUE",
			Prefix:      "DPl01",
			ExpiredDate: "2022-04-06",
			Data:        data,
		})
		repeatedCount += 1
	}
	return repeatedCount
}
func TestCreateTransaction(t *testing.T) {
	repeatedCount := Loop(10000)
	expected := 10000

	if repeatedCount != expected {
		t.Errorf("expected %q but got %q", expected, repeatedCount)
	}
}

func BenchmarkCreateTransaction(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
