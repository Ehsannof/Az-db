package models

func IsAdmin(email string) (bool) {
	var err error
	var flag = false
	
	users, err := GetAllUsers()
	if err != nil{
		return flag
	}

	personel, err := GetAllPersonel()
	if err != nil{
		return flag
	}

	for _, p := range personel {
		if p.Email == email {
			for _, u := range users{
				if u.Email == email{
					flag = true
				}
			}
		}
	}

	return flag
}
