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
	})

	allPort = append(allPort, PortfolioData{
		Year:         "2019-2020",
		HeaderImage:  "https://firebasestorage.googleapis.com/v0/b/discord-2a9c9.appspot.com/o/img2.PNG?alt=media&token=c07079f6-6379-4d99-9a45-596feed77c3a",
		Name:         "Bespoke Cake Editor",
		Technologies: []string{"VUE JS", "NODE JS", "SVG"},
		Slug:         "bespoke-cake-editor",
		SampleLink:   "https://upbeat-curie-5573c6.netlify.app/?test=dd",
		IsSample:     true,
	})

	allPort = append(allPort, PortfolioData{
		Year:         "2018-2019",
		HeaderImage:  "https://firebasestorage.googleapis.com/v0/b/discord-2a9c9.appspot.com/o/Screenshot%20from%202021-08-16%2007-45-18.png?alt=media&token=3f5e9861-41bc-4c86-ad4c-7d48ca436a08",
		Name:         "(Bespoke) HR Employee Time Track",
		Technologies: []string{"FIREBASE/FIRESTORE", "Vue JS"},
		Slug:         "hr-employee-time-track-software",
		SampleLink:   "https://freeie-d859c.firebaseapp.com/#/recent",
		IsSample:     true,
	})

	allPort = append(allPort, PortfolioData{
		Year:         "2018",
		HeaderImage:  "https://firebasestorage.googleapis.com/v0/b/discord-2a9c9.appspot.com/o/IMG.PNG?alt=media&token=945e5b4f-6d8f-4189-b98b-2157fe38625e",
		Name:         "(Bespoke) E-Commerce Website",
		Technologies: []string{"Shopify", "Laravel/PHP", "Vue JS"},
		Slug:         "bespoke-ecommerce-website",
		SampleLink:   "https://colorwayadv.web.app/#/",
		IsSample:     true,
	})

	allPort = append(allPort, PortfolioData{
		Year:         "2018-2019",
		HeaderImage:  "https://firebasestorage.googleapis.com/v0/b/discord-2a9c9.appspot.com/o/img4.PNG?alt=media&token=cfa55a8b-4a8e-45a6-9d56-eb6adfbc729c",
		Name:         "(Bespoke) Factory Fulfillment Software",
		Technologies: []string{"NODE JS", "PHP", "Vue JS"},
		Slug:         "order-management-factory-software",
		SampleLink:   "",
		IsSample:     false,
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
