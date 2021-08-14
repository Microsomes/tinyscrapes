package helpers

type SkillSample struct {
	Content string
	IsDemo  bool
	Url     string
	Year    string
	Type    string
}
type Skills struct {
	Type            string
	Name            string
	AmountOfWebsite int
	ExperienceYears int
	Url             string
	ImageURL        string
	Samples         []SkillSample
}

type Wrap struct {
	Data []Skills
	Type string
}

type ToReturn struct {
	Status    string
	AllSkills []Wrap
}

func GetAllSkills(c chan []Skills) {
	//allskills slice
	var allSkills []Skills
	allSkills = append(allSkills, Skills{
		Type:            "BACKEND",
		Name:            "NODE JS",
		AmountOfWebsite: 4,
		ExperienceYears: 3,
		Url:             "https://nodejs.org/en/",
		Samples: []SkillSample{
			SkillSample{
				Content: "Rest API for colorwayadvertising, Handled Buying (SagePay), Adding Products Etc",
				Url:     "https://colorwayadvertising.com/",
				IsDemo:  true,
				Year:    "2017",
				Type:    "commercial",
			},
			SkillSample{
				Content: "Web Scrapping, Used Pupeteer for browser testing, and web scrapping",
				Url:     "https://medium.com/me/stats/post/21ab2b627dc9",
				IsDemo:  true,
				Year:    "2020",
				Type:    "commercial",
			},
			SkillSample{
				Content: "Shopify API, Aggregate data for analysis",
				IsDemo:  false,
				Year:    "2020",
				Type:    "commercial",
			},
		},
	})
	allSkills = append(allSkills, Skills{
		Type:            "BACKEND",
		Name:            "GOLANG",
		AmountOfWebsite: 1,
		ExperienceYears: 1,
		Url:             "https://golang.org/",
		Samples: []SkillSample{
			SkillSample{
				Content: "This website, TayyabJaved.com was created using Golang, Specifically using the html/template package",
				Url:     "https://tayyabjaved.com",
				IsDemo:  true,
				Year:    "2021",
				Type:    "personal",
			},
		},
	})
	allSkills = append(allSkills, Skills{
		Type:            "BACKEND",
		Name:            "PHP/Laravel/Magento",
		AmountOfWebsite: 3,
		ExperienceYears: 4,
		Url:             "https://laravel.com/",
		Samples: []SkillSample{
			SkillSample{
				Content: "This is a laravel project, i was told to do a code test a while ago",
				Url:     "http://138.68.163.183/",
				Type:    "personal",
			},
		},
	})

	allSkills = append(allSkills, Skills{
		Type:            "MOBILE",
		Name:            "DART/Flutter",
		AmountOfWebsite: 3,
		ExperienceYears: 4,
		Url:             "https://flutter.dev/",
	})

	allSkills = append(allSkills, Skills{
		Type:            "MOBILE",
		Name:            "JAVA",
		AmountOfWebsite: 3,
		ExperienceYears: 4,
		Url:             "https://www.java.com/en/",
	})

	allSkills = append(allSkills, Skills{
		Type:            "MOBILE",
		Name:            "KOTLIN",
		AmountOfWebsite: 3,
		ExperienceYears: 4,
		Url:             "https://kotlinlang.org/",
	})

	allSkills = append(allSkills, Skills{
		Type:            "DATABASE",
		Name:            "MYSQL",
		AmountOfWebsite: 4,
		ExperienceYears: 5,
		Url:             "https://www.mysql.com/",
	})

	allSkills = append(allSkills, Skills{
		Type:            "DATABASE",
		Name:            "FIREBASE/FIRESTORE",
		AmountOfWebsite: 3,
		ExperienceYears: 3,
		Url:             "https://firebase.google.com/",
	})

	allSkills = append(allSkills, Skills{
		Type:            "DATABASE",
		Name:            "MONGODB",
		AmountOfWebsite: 3,
		ExperienceYears: 3,
		Url:             "https://www.mongodb.com/",
	})

	//postgresql

	allSkills = append(allSkills, Skills{
		Type:            "DATABASE",
		Name:            "POSTGRESQL",
		AmountOfWebsite: 3,
		ExperienceYears: 2,
		Url:             "https://www.postgresql.org/",
	})

	allSkills = append(allSkills, Skills{
		Type:            "FRONTEND",
		Name:            "VUE JS",
		AmountOfWebsite: 5,
		ExperienceYears: 4,
		Url:             "https://vuejs.org/",
		Samples: []SkillSample{
			SkillSample{
				Content: "Bespoke Cake Editor",
				Url:     "https://upbeat-curie-5573c6.netlify.app/?test=dd",
				Type:    "commercial",
			},
		},
	})

	allSkills = append(allSkills, Skills{
		Type:            "FRONTEND",
		Name:            "React",
		AmountOfWebsite: 2,
		ExperienceYears: 2,
		Url:             "https://reactjs.org/",
		Samples:         []SkillSample{},
	})

	allSkills = append(allSkills, Skills{
		Type:            "FRONTEND",
		Name:            "Liquid Shopify",
		AmountOfWebsite: 2,
		ExperienceYears: 2,
		Url:             "https://reactjs.org/",
		Samples:         []SkillSample{},
	})

	allSkills = append(allSkills, Skills{
		Type:            "FRONTEND",
		Name:            "Handle Bars/ Razor Templates",
		AmountOfWebsite: 2,
		ExperienceYears: 2,
		Url:             "https://reactjs.org/",
		Samples:         []SkillSample{},
	})

	allSkills = append(allSkills, Skills{
		Type:            "DEVOPS",
		Name:            "Git",
		AmountOfWebsite: 5,
		ExperienceYears: 5,
		Url:             "https://reactjs.org/",
		Samples:         []SkillSample{},
	})

	allSkills = append(allSkills, Skills{
		Type:            "DEVOPS",
		Name:            "Docker",
		AmountOfWebsite: 5,
		ExperienceYears: 5,
		Url:             "https://reactjs.org/",
		Samples:         []SkillSample{},
	})

	allSkills = append(allSkills, Skills{
		Type:            "DEVOPS",
		Name:            "AWS/GCP",
		AmountOfWebsite: 5,
		ExperienceYears: 5,
		Url:             "https://reactjs.org/",
		Samples:         []SkillSample{},
	})

	allSkills = append(allSkills, Skills{
		Type:            "ECOMM",
		Name:            "Shopify/App/Theme",
		AmountOfWebsite: 5,
		ExperienceYears: 2,
		Url:             "https://www.shopify.co.uk/",
		Samples:         []SkillSample{},
	})
	allSkills = append(allSkills, Skills{
		Type:            "ECOMM",
		Name:            "Woocommerce/Wordpress",
		AmountOfWebsite: 5,
		ExperienceYears: 2,
		Url:             "https://woocommerce.com/",
		Samples:         []SkillSample{},
	})

	allSkills = append(allSkills, Skills{
		Type:            "ECOMM",
		Name:            "Magento/Theme/App Development",
		AmountOfWebsite: 5,
		ExperienceYears: 2,
		Url:             "https://magento.com/",
		Samples:         []SkillSample{},
	})

	c <- allSkills

}

func GetSkillsWithType(Type string) []Skills {
	c := make(chan []Skills)
	go GetAllSkills(c)
	x := <-c
	var newList []Skills
	for _, e := range x {
		if e.Type == Type {
			newList = append(newList, e)
		}
	}
	return newList
}

func GetSkills() ToReturn {
	frontend := GetSkillsWithType("FRONTEND")
	backend := GetSkillsWithType("BACKEND")
	database := GetSkillsWithType("DATABASE")
	devops := GetSkillsWithType("DEVOPS")
	mobile := GetSkillsWithType("MOBILE")
	ecom := GetSkillsWithType("ECOMM")

	return ToReturn{
		Status: "OK",
		AllSkills: []Wrap{
			Wrap{
				Data: frontend,
				Type: "FRONTEND",
			},
			Wrap{
				Data: backend,
				Type: "BACKEND",
			},
			Wrap{
				Data: database,
				Type: "DATABASE",
			},
			Wrap{
				Data: mobile,
				Type: "MOBILE",
			},
			Wrap{
				Data: devops,
				Type: "DEVOPS",
			},
			Wrap{
				Data: ecom,
				Type: "E-commerce",
			},
		},
	}

}
