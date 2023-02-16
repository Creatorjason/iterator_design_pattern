package main



type userCollection struct{
	users []*user
}

func (uc *userCollection) createIterator() iterator{
	return &userIterator{
		users : uc.users,
	}
}