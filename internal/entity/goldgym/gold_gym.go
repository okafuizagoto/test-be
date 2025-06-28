package goldgym

import (
	"time"

	"gopkg.in/guregu/null.v3/zero"
)

type GetGoldUser struct {
	GoldId            int    `db:"gold_id" json:"gold_id"`
	GoldEmail         string `db:"gold_email" json:"gold_email"`
	GoldPassword      string `db:"gold_password" json:"gold_password"`
	GoldNama          string `db:"gold_nama" json:"gold_nama"`
	GoldNomorHp       string `db:"gold_nomorhp" json:"gold_nomorhp"`
	GoldNomorKartu    string `db:"gold_nomorkartu" json:"gold_nomorkartu"`
	GoldCvv           string `db:"gold_cvv" json:"gold_cvv"`
	GoldExpireddate   string `db:"gold_expireddate" json:"gold_expireddate"`
	GoldPemegangKartu string `db:"gold_namapemegangkartu" json:"gold_namapemegangkartu"`
}

type GetGoldUsers struct {
	GoldId            int    `db:"gold_id" json:"gold_id"`
	GoldEmail         string `db:"gold_email" json:"gold_email"`
	GoldPassword      string `db:"gold_password" json:"gold_password"`
	GoldNama          string `db:"gold_nama" json:"gold_nama"`
	GoldNomorHp       string `db:"gold_nomorhp" json:"gold_nomorhp"`
	GoldNomorKartu    string `db:"gold_nomorkartu" json:"gold_nomorkartu"`
	GoldCvv           string `db:"gold_cvv" json:"gold_cvv"`
	GoldExpireddate   string `db:"gold_expireddate" json:"gold_expireddate"`
	GoldPemegangKartu string `db:"gold_namapemegangkartu" json:"gold_namapemegangkartu"`
	GoldOTP           string `db:"gold_otp" json:"gold_otp"`
}

type GetGoldUserss struct {
	GoldId                  int         `db:"gold_id" json:"gold_id"`
	GoldEmail               string      `db:"gold_email" json:"gold_email"`
	GoldPassword            string      `db:"gold_password" json:"gold_password"`
	GoldNama                string      `db:"gold_nama" json:"gold_nama"`
	GoldNomorHp             string      `db:"gold_nomorhp" json:"gold_nomorhp"`
	GoldNomorKartu          string      `db:"gold_nomorkartu" json:"gold_nomorkartu"`
	GoldCvv                 string      `db:"gold_cvv" json:"gold_cvv"`
	GoldExpireddate         string      `db:"gold_expireddate" json:"gold_expireddate"`
	GoldPemegangKartu       string      `db:"gold_namapemegangkartu" json:"gold_namapemegangkartu"`
	GoldValidasiYN          string      `db:"gold_validasiyn" json:"gold_validasiyn"`
	GoldToken               zero.String `db:"gold_token" json:"gold_token"`
	GoldOtp                 string      `db:"gold_otp" json:"gold_otp"`
	GoldUpdatedBy           string      `db:"gold_updated_by" json:"gold_updated_by"`
	GoldUpdatedAt           string      `db:"gold_updated_at" json:"gold_updated_at"`
	GoldLastLogin           string      `db:"gold_last_login" json:"gold_last_login"`
	GoldLastLoginHost       string      `db:"gold_last_login_host" json:"gold_last_login_host"`
	GoldForceChangePassword int         `db:"gold_force_change_password" json:"gold_force_change_password"`
}

type LoginUser struct {
	GoldNama          string `db:"gold_nama" json:"gold_nama"`
	GoldNomorHp       string `db:"gold_nomorhp" json:"gold_nomorhp"`
	GoldNomorKartu    string `db:"gold_nomorkartu" json:"gold_nomorkartu"`
	GoldCvv           string `db:"gold_cvv" json:"gold_cvv"`
	GoldExpireddate   string `db:"gold_expireddate" json:"gold_expireddate"`
	GoldPemegangKartu string `db:"gold_namapemegangkartu" json:"gold_namapemegangkartu"`
	GoldToken         string `json:"gold_token"`
}

type LogUser struct {
	GoldEmail    string `json:"gold_email"`
	GoldPassword string `json:"gold_password"`
}

type LoginToken struct {
	GoldToken string `db:"gold_token" json:"gold_token"`
}

type LoginTokenDataPeserta struct {
	GoldToken string `db:"gold_token" json:"gold_token"`
	GoldEmail string `db:"gold_email" json:"gold_email"`
}

type Subscription struct {
	// GoldMenuId          int     `db:"gold_menuid" json:"gold_menuid"`
	GoldNamaPaket       string  `db:"gold_namapaket" json:"gold_namapaket"`
	GoldNamaLayanan     string  `db:"gold_namalayanan" json:"gold_namalayanan"`
	GoldHarga           float64 `db:"gold_harga" json:"gold_harga"`
	GoldJadwal          string  `db:"gold_jadwal" json:"gold_jadwal"`
	GoldListLatihan     string  `db:"gold_listlatihan" json:"gold_listlatihan"`
	GoldJumlahpertemuan int     `db:"gold_jumlahpertemuan" json:"gold_jumlahpertemuan"`
	GoldDurasi          int     `db:"gold_durasi" json:"gold_durasi"`
}

type SubscriptionAll struct {
	GoldEmail      string    `json:"gold_email"`
	GoldId         int       `db:"gold_id" json:"gold_id"`
	GoldTotalharga float64   `db:"gold_totalharga" json:"gold_totalharga"`
	GoldLastupdate time.Time `db:"gold_lastupdate" json:"gold_lastupdate"`
	// GoldMenuId int `db:"gold_menuid" json:"gold_menuid"`
	// GoldNamaPaket   string  `db:"gold_namapaket" json:"gold_namapaket"`
	// GoldNamaLayanan string  `db:"gold_namalayanan" json:"gold_namalayanan"`
	// GoldHarga       float64 `db:"gold_harga" json:"gold_harga"`
}

type SubscriptionDetail struct {
	GoldId              int     `db:"gold_id" json:"gold_id"`
	GoldMenuId          int     `db:"gold_menuid" json:"gold_menuid"`
	GoldNamaPaket       string  `db:"gold_namapaket" json:"gold_namapaket"`
	GoldNamaLayanan     string  `db:"gold_namalayanan" json:"gold_namalayanan"`
	GoldHarga           float64 `db:"gold_harga" json:"gold_harga"`
	GoldJadwal          string  `db:"gold_jadwal" json:"gold_jadwal"`
	GoldListLatihan     string  `db:"gold_listlatihan" json:"gold_listlatihan"`
	GoldJumlahpertemuan int     `db:"gold_jumlahpertemuan" json:"gold_jumlahpertemuan"`
	GoldDurasi          int     `db:"gold_durasi" json:"gold_durasi"`
	GoldStatuslangganan string  `db:"gold_statuslangganan" json:"gold_statuslangganan"`
}

type InsertSubsAll struct {
	HeaderData SubscriptionAll      `json:"header"`
	DetailData []SubscriptionDetail `json:"detail"`
}

// type DeleteSubsHeader struct {
// 	GoldId int `db:"gold_id" json:"gold_id"`
// 	// GoldMenuId int `db:"gold_menuid" json:"gold_menuid"`
// }

type DeleteSubs struct {
	GoldId     int `db:"gold_id" json:"gold_id"`
	GoldMenuId int `db:"gold_menuid" json:"gold_menuid"`
}

type UpdateSubs struct {
	GoldJumlahpertemuan int `db:"gold_jumlahpertemuan" json:"gold_jumlahpertemuan"`
	GoldId              int `db:"gold_id" json:"gold_id"`
	GoldMenuId          int `db:"gold_menuid" json:"gold_menuid"`
}

type UpdatePassword struct {
	GoldPassword string `db:"gold_password" json:"gold_password"`
	GoldEmail    string `db:"gold_email" json:"gold_email"`
	GoldOTP      string `db:"gold_otp" json:"gold_otp"`
}

type UpdateNama struct {
	GoldNama  string `db:"gold_nama" json:"gold_nama"`
	GoldEmail string `db:"gold_email" json:"gold_email"`
}

type UpdateKartu struct {
	GoldNomorKartu string `db:"gold_nomorkartu" json:"gold_nomorkartu"`
	GoldCvv        string `db:"gold_cvv" json:"gold_cvv"`
	GoldEmail      string `db:"gold_email" json:"gold_email"`
}

type Logout struct {
	GoldEmail string `db:"gold_email" json:"gold_email"`
}

type GetSubsWithUser struct {
	GoldId              int         `db:"gold_id" json:"gold_id"`
	GoldMenuId          zero.String `db:"gold_menuid" json:"gold_menuid"`
	GoldEmail           string      `db:"gold_email" json:"gold_email"`
	GoldNama            string      `db:"gold_nama" json:"gold_nama"`
	GoldNomorHp         string      `db:"gold_nomorhp" json:"gold_nomorhp"`
	GoldExpireddate     string      `db:"gold_expireddate" json:"gold_expireddate"`
	GoldNamaPaket       zero.String `db:"gold_namapaket" json:"gold_namapaket"`
	GoldNamaLayanan     zero.String `db:"gold_namalayanan" json:"gold_namalayanan"`
	GoldHarga           zero.Float  `db:"gold_harga" json:"gold_harga"`
	GoldListLatihan     zero.String `db:"gold_listlatihan" json:"gold_listlatihan"`
	GoldJumlahpertemuan zero.Int    `db:"gold_jumlahpertemuan" json:"gold_jumlahpertemuan"`
	GoldDurasi          zero.Int    `db:"gold_durasi" json:"gold_durasi"`
	GoldStatuslangganan zero.String `db:"gold_statuslangganan" json:"gold_statuslangganan"`
}

type GetValidationGoldOTP struct {
	GoldOTP string `db:"gold_otp" json:"gold_otp"`
}

type UpdateValidationOTP struct {
	GoldEmail string `db:"gold_email" json:"gold_email"`
}

type SubscriptionHeader struct {
	GoldID              int         `db:"gold_id" json:"gold_id"`
	GoldTotalharga      zero.Float  `db:"gold_totalharga" json:"gold_totalharga"`
	GoldValidasiPayment string      `db:"gold_validasipayment" json:"gold_validasipayment"`
	GoldOTP             zero.String `db:"gold_otp" json:"gold_otp"`
	GoldLastupdate      zero.String `db:"gold_lastupdate" json:"gold_lastupdate"`
}

type SubscriptionHeaderPayment struct {
	// GoldID              int         `db:"gold_id" json:"gold_id"`
	GoldTotalharga      zero.Float `db:"gold_totalharga" json:"gold_totalharga"`
	GoldValidasiPayment string     `db:"gold_validasipayment" json:"gold_validasipayment"`
	// GoldOTP             zero.String `db:"gold_otp" json:"gold_otp"`
	// GoldLastupdate      zero.String `db:"gold_lastupdate" json:"gold_lastupdate"`
}

type UpdatePayment struct {
	GoldID int `db:"gold_id" json:"gold_id"`
}

// testings
type Testings struct {
	ID            string `json:"id"`
	TestingImages []byte `json:"testing"`
}
