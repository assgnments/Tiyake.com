package form

type VaildationErros map[string][]string

func (ve VaildationErros) Add (field string,message string){
	ve[field]=append(ve[field],message)
}
//Get message for the template to get the first error
func (ve VaildationErros) Get(field string) string {
	errors:=ve[field]
	if len(errors)==0 {
		return ""
	}
	return errors[0]
}