package licenses

// func LicenseActionFromActivationKey(dto *LicenseFromActivationKeyDto, query fireback.QueryDSL) (*LicenseEntity, *fireback.IError) {

// 	license := &LicenseEntity{}

// 	query2 := query

// 	query2.UniqueId = dto.ActivationKeyId
// 	query.Deep = true
// 	ak, err := activationkeys.ActivationKeyActionGetOne(query2)

// 	if err != nil {
// 		return nil, err
// 	}

// 	if ak.Used {
// 		return nil, fireback.CreateIErrorString("USED_ACTIVATION_KEY", []string{}, 403)
// 	}

// 	license.Title = ak.Plan.Title
// 	license.ValidityStartDate = time.Now()
// 	license.ValidityEndDate = time.Now().Add(time.Hour * 24 * time.Duration(ak.Plan.Duration))
// 	license.WorkspaceId = query.WorkspaceId
// 	license.UserId = query.UserId

// 	data := &LicenseConfigurationFile{}
// 	fireback.ReadYamlFile("licenses/fireback.yml", data)

// 	doc := LicenseContent{
// 		Email:             "user@domain.com",
// 		ValidityStartDate: license.ValidityStartDate,
// 		ValidityEndDate:   license.ValidityEndDate,
// 		WorkspaceId:       license.WorkspaceId,
// 		MachineId:         dto.MachineId,
// 	}

// 	licenseE, err2 := GenerateLicense(doc, data.PrivateKey)
// 	license.SignedLicense = licenseE

// 	if err2 != nil {
// 		fmt.Println("Error on license", err)
// 	} else {
// 		fmt.Println(licenseE)
// 	}

// 	license, err = LicenseActionCreate(license, query)

// 	if err != nil {
// 		return nil, err
// 	}

// 	_, err3 := activationkeys.ActivationKeyActionUpdate(query2, &activationkeys.ActivationKeyEntity{Used: true})

// 	if err3 != nil {
// 		return nil, err
// 	}

// 	event.MustFire(LICENSE_EVENT_CREATED, event.M{
// 		"entity":   license,
// 		"target":   "workspace",
// 		"unqiueId": query.WorkspaceId,
// 	})

// 	return license, nil
// }
