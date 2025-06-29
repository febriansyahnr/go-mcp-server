package secret_reader

func LoadSecrets(acquirers []string) map[string]SecretReader {
	secretAcquirers := make(map[string]SecretReader)

	for _, acquirer := range acquirers {
		switch acquirer {
		case "snap_core":
			secretAcquirers["snap_core"] = New("snap_core")
		case "permata":
			secretAcquirers["permata"] = New("permata")
		case "aspi":
			secretAcquirers["aspi"] = New("aspi")
		case "bnc":
			secretAcquirers["bnc"] = New("bnc")
		case "qa_asymmetric":
			secretAcquirers["qa_asymmetric"] = New("qa_asymmetric")
		case "bri_qris":
			secretAcquirers["bri_qris"] = New("bri_qris")
		case "mandiri_central":
			secretAcquirers["mandiri_central"] = New("mandiri_central")
		}
	}

	return secretAcquirers
}
