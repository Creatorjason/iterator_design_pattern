package main

type userIterator struct{
	index int
	users []*user
}

func (ui *userIterator) hasNext() bool{
	return ui.index >= len(ui.users)
}


func (ui *userIterator) getNext() *user{
	if ui.hasNext(){
		user := ui.users[ui.index]
		ui.index++
		return user
	}
	return nil
}