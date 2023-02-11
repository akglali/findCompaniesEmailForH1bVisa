package main

import "findCompaniesEmailForH1bVisa/sendEmailToFoundCompanies"

func main() {
	// **** h1bgrader.com version
	//first getting the companiesName name from h1b grader website for software developer position
	//companiesName.GetCompaniesName()
	//then find the companiesName website according to their name
	//companiesWebsites.FindWebsites()
	//time to find companies email if it is existed and save them to use it.
	//companiesEmail.FindEmail()
	//last step send the all found emails.
	//sendEmailToFoundCompanies.SendEmail()

	// **** Google Map Search Version
	//get all the companies information
	//googleMapsSearch.getCompaniesNameByGoogleSearch()

	//find all the companies website by their name
	//googleMapsSearch.FindWebsites()

	//find companies email from their websites
	//googleMapsSearch.FindEmailGoogleMap()

	//send emails
	sendEmailToFoundCompanies.SendEmails()
}
