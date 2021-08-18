package helpers

type PortfolioContent struct {
	Title   string
	Subline string
}

type PortfolioData struct {
	HeaderImage  string
	Name         string
	Technologies []string
	Slug         string
	Extra        PortfolioContent
	SampleLink   string
	IsSample     bool
	Year         string
	Subline      string
	Screenshots  [3]string
	IsScreenShot bool
}

func GetPort() []PortfolioData {
	allPort := []PortfolioData{}

	allPort = append(allPort, PortfolioData{
		Year:         "2019-2020",
		HeaderImage:  "https://firebasestorage.googleapis.com/v0/b/discord-2a9c9.appspot.com/o/img3.PNG?alt=media&token=4aa0b816-7826-4047-8766-0b7cb1ba555f",
		Name:         "Historic Newspapers",
		Technologies: []string{"Wordpress", "WooCommerce", "PHP"},
		Slug:         "historic-newspapers",
		SampleLink:   "https://www.historic-newspapers.co.uk/",
		IsSample:     true,
		Subline:      "Taking on a badly maintained/written code base, and fixing enormous amounts of bugs, to get the company back to selling...",
		Screenshots: [3]string{
			"https://firebasestorage.googleapis.com/v0/b/discord-2a9c9.appspot.com/o/historicnewspapers%2FScreenshot%20from%202021-08-18%2012-10-56.png?alt=media&token=0eeb274e-3f35-4039-8e06-d230cb044a7e",
			"https://firebasestorage.googleapis.com/v0/b/discord-2a9c9.appspot.com/o/historicnewspapers%2FScreenshot%20from%202021-08-18%2012-12-22.png?alt=media&token=2dc0b73e-2240-4604-bb59-c6348551bb61",
			"https://firebasestorage.googleapis.com/v0/b/discord-2a9c9.appspot.com/o/historicnewspapers%2FScreenshot%20from%202021-08-18%2012-18-18.png?alt=media&token=db14e5c5-7f78-43e5-a36a-1f355a233460",
		},
		IsScreenShot: true,
	})

	allPort = append(allPort, PortfolioData{
		Year:         "2019-2020",
		HeaderImage:  "https://firebasestorage.googleapis.com/v0/b/discord-2a9c9.appspot.com/o/img2.PNG?alt=media&token=c07079f6-6379-4d99-9a45-596feed77c3a",
		Name:         "Bespoke Cake Editor",
		Technologies: []string{"VUE JS", "NODE JS", "SVG"},
		Slug:         "bespoke-cake-editor",
		SampleLink:   "https://upbeat-curie-5573c6.netlify.app/?test=dd",
		IsSample:     true,
		Subline:      "Creating,Maintaining and improving the customer statisfaction by creating a fully bespoke edit your own cake feature",
		Screenshots: [3]string{
			"https://firebasestorage.googleapis.com/v0/b/discord-2a9c9.appspot.com/o/bespokecake%2FScreenshot%20from%202021-08-18%2012-20-43.png?alt=media&token=1a3f425f-dcf0-44d9-b223-0809ced7dd1d",
			"https://firebasestorage.googleapis.com/v0/b/discord-2a9c9.appspot.com/o/bespokecake%2FScreenshot%20from%202021-08-18%2012-23-34.png?alt=media&token=0510d518-96d5-428f-8ac4-9e8c11e840d2",
			"https://firebasestorage.googleapis.com/v0/b/discord-2a9c9.appspot.com/o/bespokecake%2FScreenshot%20from%202021-08-18%2012-24-05.png?alt=media&token=c214b53b-a78a-4acc-8579-9a6facbd05cd",
		},
		IsScreenShot: true,
	})

	allPort = append(allPort, PortfolioData{
		Year:         "2018-2019",
		HeaderImage:  "https://firebasestorage.googleapis.com/v0/b/discord-2a9c9.appspot.com/o/Screenshot%20from%202021-08-16%2007-45-18.png?alt=media&token=3f5e9861-41bc-4c86-ad4c-7d48ca436a08",
		Name:         "(Bespoke) HR Employee Time Track",
		Technologies: []string{"FIREBASE/FIRESTORE", "Vue JS"},
		Slug:         "hr-employee-time-track-software",
		SampleLink:   "https://freeie-d859c.firebaseapp.com/#/recent",
		IsSample:     true,
		Subline:      "Reducing errors, and increasting time efficiency of logging and calculating employee work hours. ",
		Screenshots: [3]string{
			"https://firebasestorage.googleapis.com/v0/b/discord-2a9c9.appspot.com/o/Screenshot%20from%202021-08-16%2007-45-18.png?alt=media&token=3f5e9861-41bc-4c86-ad4c-7d48ca436a08",
			"https://firebasestorage.googleapis.com/v0/b/discord-2a9c9.appspot.com/o/timesheet%2FScreenshot%20from%202021-08-18%2012-29-23.png?alt=media&token=91d0acd7-a770-4442-b044-750cb485c563",
			"https://firebasestorage.googleapis.com/v0/b/discord-2a9c9.appspot.com/o/timesheet%2FScreenshot%20from%202021-08-18%2012-30-53.png?alt=media&token=78d2a526-1f70-45af-ba1c-414ce39998f5",
		},
		IsScreenShot: true,
	})

	allPort = append(allPort, PortfolioData{
		Year:         "2018",
		HeaderImage:  "https://firebasestorage.googleapis.com/v0/b/discord-2a9c9.appspot.com/o/IMG.PNG?alt=media&token=945e5b4f-6d8f-4189-b98b-2157fe38625e",
		Name:         "(Bespoke) E-Commerce Website",
		Technologies: []string{"Shopify", "Laravel/PHP", "Vue JS"},
		Slug:         "bespoke-ecommerce-website",
		SampleLink:   "https://colorwayadv.web.app/#/",
		IsSample:     true,
		Subline:      "Re-designing their main E-Commerce website, increasting sales and functionality",
		Screenshots: [3]string{
			"https://firebasestorage.googleapis.com/v0/b/discord-2a9c9.appspot.com/o/bespokecom%2FScreenshot%20from%202021-08-18%2012-34-31.png?alt=media&token=18ef2afd-5aad-4b72-98dd-36aee6016a47",
			"https://firebasestorage.googleapis.com/v0/b/discord-2a9c9.appspot.com/o/bespokecom%2FScreenshot%20from%202021-08-18%2013-05-35.png?alt=media&token=01c98981-8cd6-4b65-a7f7-a474e2750ac4",
			"https://firebasestorage.googleapis.com/v0/b/discord-2a9c9.appspot.com/o/bespokecom%2FScreenshot%20from%202021-08-18%2013-05-02.png?alt=media&token=f09fc4dd-c111-4224-9cbf-d00206fedcc9",
		},
		IsScreenShot: true,
	})

	allPort = append(allPort, PortfolioData{
		Year:         "2018-2019",
		HeaderImage:  "https://firebasestorage.googleapis.com/v0/b/discord-2a9c9.appspot.com/o/img4.PNG?alt=media&token=cfa55a8b-4a8e-45a6-9d56-eb6adfbc729c",
		Name:         "(Bespoke) Factory Fulfillment Software",
		Technologies: []string{"NODE JS", "PHP", "Vue JS"},
		Slug:         "order-management-factory-software",
		SampleLink:   "",
		IsSample:     false,
		Subline:      "Handling and tracking 1000s of daily orders, from generating shipping labels through logging each order through various stages of the factory process, this software had it all",
		IsScreenShot: false,
	})

	return allPort
}

//method
func FindBlogBySlug(slug string) PortfolioData {
	port := GetPort()
	for _, val := range port {
		if val.Slug == slug {
			return val
		}
	}
	return PortfolioData{}
}
