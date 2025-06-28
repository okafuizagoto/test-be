package goldgym

import (
	"context"
	goldEntity "gold-gym-be/internal/entity/goldgym"
	"gold-gym-be/pkg/errors"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/raja/argon2pw"
)

func (d Data) GetGoldUser(ctx context.Context) ([]goldEntity.GetGoldUser, error) {
	var (
		user  goldEntity.GetGoldUser
		users []goldEntity.GetGoldUser
		err   error
	)
	log.Println("data GetGoldUser object")
	rows, err := (*d.stmt)[getGoldUser].QueryxContext(ctx)
	if err != nil {
		return users, errors.Wrap(err, "[DATA] [GetGoldUser]")
	}
	log.Println("datagolduser", users)

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&user); err != nil {
			return users, errors.Wrap(err, "[DATA] [GetGoldUser]")
		}
		users = append(users, user)
	}
	return users, err
}

func (d Data) GetGoldUserByEmail(ctx context.Context, email string) (goldEntity.GetGoldUserss, error) {
	var (
		user goldEntity.GetGoldUserss
		err  error
	)
	log.Println("data GetGoldUserByEmail object", email)
	rows, err := (*d.stmt)[getGoldUserByEmail].QueryxContext(ctx, email)
	if err != nil {
		return user, errors.Wrap(err, "[DATA] [GetGoldUserByEmail]")
	}
	log.Println("datagolduser", user)

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&user); err != nil {
			return user, errors.Wrap(err, "[DATA] [GetGoldUserByEmail]")
		}
	}
	log.Println("datagolduser2", user)
	return user, err
}

func (d Data) GetGoldUserByEmailLogin(ctx context.Context, email string, password string) (goldEntity.LoginUser, error) {
	var (
		user goldEntity.LoginUser
		err  error
	)
	log.Println("data GetGoldUserByEmailLogin object")
	rows, err := (*d.stmt)[getGoldUserByEmailLogin].QueryxContext(ctx, email, password)
	if err != nil {
		return user, errors.Wrap(err, "[DATA] [GetGoldUserByEmailLogin]")
	}
	log.Println("dataGetGoldUserByEmailLogin", user)

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&user); err != nil {
			return user, errors.Wrap(err, "[DATA] [GetGoldUserByEmailLogin]")
		}
	}
	return user, err
}

func (d Data) InsertGoldUser(ctx context.Context, user goldEntity.GetGoldUsers) (string, error) {
	var result string
	var err error

	_, err = (*d.stmt)[insertGoldUser].ExecContext(ctx,
		user.GoldId,
		user.GoldEmail,
		user.GoldPassword,
		user.GoldNama,
		user.GoldNomorHp,
		user.GoldNomorKartu,
		user.GoldCvv,
		user.GoldExpireddate,
		user.GoldPemegangKartu,
		user.GoldOTP,
	)

	log.Println("data user object", user)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][InsertGoldUser]")
	}
	result = "Sukses"

	return result, err

}

func (d Data) GetGoldToken(ctx context.Context) (goldEntity.LoginToken, error) {
	var (
		user goldEntity.LoginToken
		err  error
	)
	log.Println("data GetGoldToken object")
	rows, err := (*d.stmt)[getGoldToken].QueryxContext(ctx)
	if err != nil {
		return user, errors.Wrap(err, "[DATA] [GetGoldToken]")
	}
	log.Println("dataGetGoldToken", user)

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&user); err != nil {
			return user, errors.Wrap(err, "[DATA] [GetGoldToken]")
		}
	}
	return user, err
}

func (d Data) UpdateGoldToken(ctx context.Context, user goldEntity.LoginTokenDataPeserta) error {
	var err error

	_, err = (*d.stmt)[updateGoldToken].ExecContext(ctx,
		user.GoldToken,
		user.GoldEmail,
	)

	log.Println("data user object", user)

	if err != nil {
		return errors.Wrap(err, "[DATA][InsertGoldUser]")
	}

	return err

}

func (d Data) GetAllSubscription(ctx context.Context) ([]goldEntity.Subscription, error) {
	var (
		user  goldEntity.Subscription
		users []goldEntity.Subscription
		err   error
	)
	log.Println("data GetAllSubscription object")
	rows, err := (*d.stmt)[getAllSubscription].QueryxContext(ctx)
	if err != nil {
		return users, errors.Wrap(err, "[DATA] [GetAllSubscription]")
	}
	log.Println("dataGetAllSubscription", users)

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&user); err != nil {
			return users, errors.Wrap(err, "[DATA] [GetAllSubscription]")
		}
		users = append(users, user)
	}
	return users, err
}

func (d Data) InsertSubscription(ctx context.Context, user goldEntity.SubscriptionAll) error {

	var err error

	_, err = (*d.stmt)[insertSubscription].ExecContext(ctx,
		user.GoldId,
		user.GoldTotalharga,
		// user.GoldMenuId,
		// user.GoldNamaPaket,
		// user.GoldNamaLayanan,
		// user.GoldHarga,
	)

	log.Println("data user object", user)

	if err != nil {
		return errors.Wrap(err, "[DATA][InsertSubscription]")
	}

	return err

}

func (d Data) InsertSubscriptionDetail(ctx context.Context, user goldEntity.SubscriptionDetail) error {

	var err error

	_, err = (*d.stmt)[insertSubscriptionDetail].ExecContext(ctx,
		user.GoldId,
		user.GoldMenuId,
		user.GoldNamaPaket,
		user.GoldNamaLayanan,
		user.GoldHarga,
		user.GoldJadwal,
		user.GoldListLatihan,
		user.GoldJumlahpertemuan,
		user.GoldDurasi,
		user.GoldStatuslangganan,
	)

	log.Println("data user object", user)

	if err != nil {
		return errors.Wrap(err, "[DATA][InsertSubscriptionDetail]")
	}
	// result = "Sukses"

	return err

}

// func (d Data) DeleteSubscriptionHeader(ctx context.Context, user goldEntity.DeleteSubsHeader) error {

// 	var err error

// 	_, err = (*d.stmt)[deleteSubscriptionHeader].ExecContext(ctx,
// 		user.GoldId,
// 		// user.GoldMenuId,
// 	)

// 	log.Println("data user object", user)

// 	if err != nil {
// 		return errors.Wrap(err, "[DATA][InsertSubscription]")
// 	}

// 	return err

// }

func (d Data) DeleteSubscriptionDetail(ctx context.Context, user goldEntity.DeleteSubs) error {

	var err error

	_, err = (*d.stmt)[deleteSubscriptionDetail].ExecContext(ctx,
		user.GoldId,
		user.GoldMenuId,
	)

	log.Println("data user object", user)

	if err != nil {
		return errors.Wrap(err, "[DATA][InsertSubscription]")
	}

	return err

}

func (d Data) UpdateSubscriptionDetail(ctx context.Context, user goldEntity.UpdateSubs) error {

	var err error

	_, err = (*d.stmt)[updateSubscriptionDetail].ExecContext(ctx,
		user.GoldJumlahpertemuan,
		user.GoldId,
		user.GoldMenuId,
	)

	log.Println("data user object", user)

	if err != nil {
		return errors.Wrap(err, "[DATA][InsertSubscription]")
	}

	return err

}

func (d Data) UpdateDataPeserta(ctx context.Context, user goldEntity.UpdatePassword) error {

	var err error

	hashedPassword, err := argon2pw.GenerateSaltedHash(user.GoldPassword)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][CreateUser]")
	}

	_, err = (*d.stmt)[updateDataPeserta].ExecContext(ctx,
		hashedPassword,
		user.GoldEmail,
		user.GoldOTP,
	)

	log.Println("data user object", user)

	if err != nil {
		return errors.Wrap(err, "[DATA][InsertSubscription]")
	}

	return err

}

func (d Data) UpdateNama(ctx context.Context, user goldEntity.UpdateNama) error {

	var err error

	_, err = (*d.stmt)[updateNama].ExecContext(ctx,
		user.GoldNama,
		user.GoldEmail,
	)

	log.Println("data user object", user)

	if err != nil {
		return errors.Wrap(err, "[DATA][InsertSubscription]")
	}

	return err

}

func (d Data) UpdateKartu(ctx context.Context, user goldEntity.UpdateKartu) error {

	var err error

	_, err = (*d.stmt)[updateKartu].ExecContext(ctx,
		user.GoldNomorKartu,
		user.GoldCvv,
		user.GoldEmail,
	)

	log.Println("data user object", user)

	if err != nil {
		return errors.Wrap(err, "[DATA][InsertSubscription]")
	}

	return err

}

func (d Data) Logout(ctx context.Context, user goldEntity.Logout) error {

	var err error

	_, err = (*d.stmt)[logout].ExecContext(ctx,
		user.GoldEmail,
	)

	log.Println("data user object", user)

	if err != nil {
		return errors.Wrap(err, "[DATA][InsertSubscription]")
	}

	return err

}

func (d Data) GetSubsWithUser(ctx context.Context) ([]goldEntity.GetSubsWithUser, error) {
	var (
		user  goldEntity.GetSubsWithUser
		users []goldEntity.GetSubsWithUser
		err   error
	)
	log.Println("data GetGoldUser object")
	rows, err := (*d.stmt)[getSubsWithUser].QueryxContext(ctx)
	if err != nil {
		return users, errors.Wrap(err, "[DATA] [GetGoldUser]")
	}
	log.Println("datagolduser", users)

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&user); err != nil {
			return users, errors.Wrap(err, "[DATA] [GetGoldUser]")
		}
		users = append(users, user)
	}
	return users, err
}

func (d Data) GetValidationGoldOTP(ctx context.Context, otp string) (goldEntity.GetValidationGoldOTP, error) {
	var (
		users goldEntity.GetValidationGoldOTP
		// users []goldEntity.GetSubsWithUser
		err error
	)
	log.Println("data GetGoldUser object")
	rows, err := (*d.stmt)[getValidationGoldOTP].QueryxContext(ctx, otp)
	if err != nil {
		return users, errors.Wrap(err, "[DATA] [GetGoldUser]")
	}
	log.Println("datagolduser", users)

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&users); err != nil {
			return users, errors.Wrap(err, "[DATA] [GetGoldUser]")
		}
		// users = append(users, user)
	}
	return users, err
}

func (d Data) UpdateValidationOTP(ctx context.Context, email string) error {
	var err error

	_, err = (*d.stmt)[updateValidationOTP].ExecContext(ctx, email)

	// log.Println("data user object", user)

	if err != nil {
		return errors.Wrap(err, "[DATA][InsertGoldUser]")
	}

	return err

}

func (d Data) UpdateOtpIsNull(ctx context.Context, email string) error {
	var err error

	_, err = (*d.stmt)[updateOtpIsNull].ExecContext(ctx, email)

	// log.Println("data user object", user)

	if err != nil {
		return errors.Wrap(err, "[DATA][InsertGoldUser]")
	}

	return err

}

func (d Data) UpdateOTP(ctx context.Context, otp string, email string) error {
	var err error

	_, err = (*d.stmt)[updateOTP].ExecContext(ctx, otp, email)

	// log.Println("data user object", user)

	if err != nil {
		return errors.Wrap(err, "[DATA][UpdateOTP]")
	}

	return err

}

func (d Data) GetOneSubscription(ctx context.Context, menuid int) (goldEntity.Subscription, error) {
	var (
		// user  goldEntity.Subscription
		users goldEntity.Subscription
		err   error
	)
	log.Println("data GetOneSubscription object")
	rows, err := (*d.stmt)[getOneSubscription].QueryxContext(ctx, menuid)
	if err != nil {
		return users, errors.Wrap(err, "[DATA] [GetOneSubscription]")
	}
	log.Println("dataGetOneSubscription", users)

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&users); err != nil {
			return users, errors.Wrap(err, "[DATA] [GetOneSubscription]")
		}
		// users = append(users, user)
	}
	return users, err
}

func (d Data) UpdateOTPSubscription(ctx context.Context, otp string, id int) error {
	var err error

	_, err = (*d.stmt)[updateOTPSubscription].ExecContext(ctx, otp, id)

	// log.Println("data user object", user)

	if err != nil {
		return errors.Wrap(err, "[DATA][UpdateOTP]")
	}

	return err

}

func (d Data) BulkInsertSubscriptionDetail(ctx context.Context, user []goldEntity.SubscriptionDetail) error {

	for _, v := range user {
		query, args, err := sqlx.In(qInsertSubscriptionDetail,
			v.GoldId, v.GoldMenuId, v.GoldNamaPaket, v.GoldNamaLayanan, v.GoldHarga, v.GoldJadwal, v.GoldListLatihan, v.GoldJumlahpertemuan, v.GoldDurasi, v.GoldStatuslangganan)
		if err != nil {
			return errors.Wrap(err, "[DATA][BulkInsertSubscriptionDetail]")
		}
		_, err = d.db.ExecContext(ctx, query, args...)
		if err != nil {
			return errors.Wrap(err, "[DATA][BulkInsertSubscriptionDetail]")
		}
	}
	return nil
}

func (d Data) GetSubscriptionHeader(ctx context.Context, id int) (goldEntity.SubscriptionHeader, error) {
	var (
		// user  goldEntity.SubscriptionHeader
		users goldEntity.SubscriptionHeader
		err   error
	)
	log.Println("data GetGoldUser object")
	rows, err := (*d.stmt)[getSubscriptionHeader].QueryxContext(ctx, id)
	if err != nil {
		return users, errors.Wrap(err, "[DATA] [GetSubscriptionHeader]")
	}
	log.Println("datagolduser", users)

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&users); err != nil {
			return users, errors.Wrap(err, "[DATA] [GetSubscriptionHeader]")
		}
		// users = append(users, user)
	}
	return users, err
}

func (d Data) UpdateValidasiPaymentHeader(ctx context.Context, updatePayment goldEntity.UpdatePayment) error {
	var err error

	_, err = (*d.stmt)[updateValidasiPaymentHeader].ExecContext(ctx, updatePayment.GoldID)

	// log.Println("data user object", user)

	if err != nil {
		return errors.Wrap(err, "[DATA][UpdateValidasiPaymentHeader]")
	}

	return err

}

func (d Data) UpdateValidasiPaymentDetail(ctx context.Context, updatePayment goldEntity.UpdatePayment) error {
	var err error

	_, err = (*d.stmt)[updateValidasiPaymentDetail].ExecContext(ctx, updatePayment.GoldID)

	// log.Println("data user object", user)

	if err != nil {
		return errors.Wrap(err, "[DATA][UpdateValidasiPaymentHeader]")
	}

	return err

}

func (d Data) GetSubscriptionHeaderTotalHarga(ctx context.Context, id int) (goldEntity.SubscriptionHeaderPayment, error) {
	var (
		// user  goldEntity.SubscriptionHeader
		users goldEntity.SubscriptionHeaderPayment
		err   error
	)
	log.Println("data GetGoldUser object")
	rows, err := (*d.stmt)[getSubscriptionHeaderTotalHarga].QueryxContext(ctx, id)
	if err != nil {
		return users, errors.Wrap(err, "[DATA] [GetSubscriptionHeader]")
	}
	log.Println("datagolduser", users)

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&users); err != nil {
			return users, errors.Wrap(err, "[DATA] [GetSubscriptionHeader]")
		}
		// users = append(users, user)
	}
	return users, err
}

// func (d Data) GetPasswordByUser(ctx context.Context, _user string) ([]goldEntity.GetGoldUser, error) {
// 	var (
// 		user  goldEntity.GetGoldUser
// 		users []goldEntity.GetGoldUser
// 		err   error
// 	)
// 	log.Println("data GetPasswordByUser object")
// 	rows, err := (*d.stmt)[getPasswordByUser].QueryxContext(ctx, _user)
// 	if err != nil {
// 		return users, errors.Wrap(err, "[DATA] [GetPasswordByUser]")
// 	}
// 	log.Println("datagolduser", users)

// 	defer rows.Close()

// 	for rows.Next() {
// 		if err = rows.StructScan(&user); err != nil {
// 			return users, errors.Wrap(err, "[DATA] [GetPasswordByUser]")
// 		}
// 		users = append(users, user)
// 	}
// 	return users, err
// }

func (d Data) GetPasswordByUser(ctx context.Context, _user string) (string, error) {
	password := ""
	if err := (*d.stmt)[getPasswordByUser].QueryRowxContext(ctx, _user).Scan(&password); err != nil {
		return "", errors.Wrap(err, "[DATA][CheckPassword]")
	}

	return password, nil
}

func (d Data) UpdateLastLogin(ctx context.Context, _user goldEntity.GetGoldUserss) error {
	_, err := (*d.stmt)[updateLastLogin].ExecContext(ctx, _user.GoldLastLoginHost, _user.GoldEmail)
	log.Printf("userUpdate %+v", _user)
	if err != nil {
		return errors.Wrap(err, "[DATA][UpdateLastLogin]")
	}

	return nil
}

func (d Data) UploadTestingImages(ctx context.Context, testing goldEntity.Testings) (string, error) {
	var result string
	var err error

	_, err = (*d.stmt)[uploadTestingImages].ExecContext(ctx,
		testing.ID,
		testing.TestingImages,
	)

	log.Println("data user object", testing)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][InsertGoldUser]")
	}
	result = "Sukses"

	return result, err

}

func (d Data) GetTestingImages(ctx context.Context, id int) ([]byte, error) {
	var image []byte
	if err := (*d.stmt)[getTestingImage].QueryRowxContext(ctx, id).Scan(&image); err != nil {
		return image, errors.Wrap(err, "[DATA][CheckPassword]")
	}

	return image, nil
}
