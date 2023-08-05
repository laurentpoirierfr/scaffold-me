
* https://earthly.dev/blog/tui-app-with-go/

```golang
func addContactForm() {
    contact := Contact{}

    form.AddInputField("First Name", "", 20, nil, func(firstName string) {
        contact.firstName = firstName
    })

    form.AddInputField("Last Name", "", 20, nil, func(lastName string) {
        contact.lastName = lastName
    })

    form.AddInputField("Email", "", 20, nil, func(email string) {
        contact.email = email
    })

    form.AddInputField("Phone", "", 20, nil, func(phone string) {
        contact.phoneNumber = phone
    })

    // states is a slice of state abbreviations. Code is in the repo. 
    form.AddDropDown("State", states, 0, func(state string, index int) {
        contact.state = state
    })

    form.AddCheckbox("Business", false, func(business bool) {
        contact.business = business
    })

    form.AddButton("Save", func() {
        contacts = append(contacts, contact)
    })
}
```