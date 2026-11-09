package main

type Contribution struct {
	Org  string `json:"org"`
	Repo string `json:"repo"`
	PR   string `json:"pr"`
	URL  string `json:"url"`
	Key  string `json:"key"`
}

type Experience struct {
	Company  string   `json:"company"`
	Role     string   `json:"role"`
	Period   string   `json:"period"`
	Location string   `json:"location"`
	Bullets  []string `json:"bullets"`
}

type Hackathon struct {
	Title       string   `json:"title"`
	Award       string   `json:"award"`
	Organizer   string   `json:"organizer"`
	Year        string   `json:"year"`
	Description string   `json:"description"`
	Tech        []string `json:"tech"`
}

type Project struct {
	Title      string   `json:"title"`
	Subtitle   string   `json:"subtitle"`
	Highlights []string `json:"highlights"`
	Tags       []string `json:"tags"`
	Github     string   `json:"github"`
	Demo       string   `json:"demo"`
}

type SkillCategory struct {
	Category string   `json:"category"`
	Skills   []string `json:"skills"`
}

type Achievement struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Certification struct {
	Name   string `json:"name"`
	Issuer string `json:"issuer"`
	Date   string `json:"date"`
}

type StackItem struct {
	Name        string `json:"name"`
	Subtitle    string `json:"subtitle"`
	Route       string `json:"route"`
	Repo        string `json:"repo"`
	Badge       string `json:"badge"`
	BadgeClass  string `json:"badge_class"`
	BorderClass string `json:"border_class"`
	Icon        string `json:"icon"`
	IsCurrent   bool   `json:"is_current"`
}

func GetPortfolioData() map[string]interface{} {
	contributions := []Contribution{
		{Org: "Apple", Repo: "swift-numerics", PR: "#355", URL: "https://github.com/apple/swift-numerics/pull/355", Key: "apple"},
		{Org: "Google", Repo: "Highway", PR: "#3511", URL: "https://github.com/google/highway/pull/3511", Key: "google"},
		{Org: "Google Cloud", Repo: "Data Validator", PR: "#1863", URL: "https://github.com/GoogleCloudPlatform/professional-services-data-validator/pull/1863", Key: "googlecloud"},
		{Org: "Microsoft", Repo: "Semantic Kernel", PR: "#14472", URL: "https://github.com/microsoft/semantic-kernel/pull/14472", Key: "microsoft"},
		{Org: "AWS", Repo: "Lambda Powertools", PR: "#8542", URL: "https://github.com/aws-powertools/powertools-lambda-python/pull/8542", Key: "aws"},
		{Org: "Meta", Repo: "StyleX", PR: "#1899", URL: "https://github.com/facebook/stylex/pull/1899", Key: "meta"},
		{Org: "NVIDIA", Repo: "cuOpt", PR: "#2068", URL: "https://github.com/NVIDIA/cuopt/pull/2068", Key: "nvidia"},
		{Org: "NASA", Repo: "ProgPy", PR: "#237", URL: "https://github.com/nasa/prog_models/pull/237", Key: "nasa"},
		{Org: "Azure", Repo: "Cosmos DB Emulator", PR: "#359", URL: "https://github.com/Azure/azure-cosmos-db-emulator-docker/pull/359", Key: "azure"},
		{Org: "Adobe", Repo: "Skills", PR: "#445", URL: "https://github.com/adobe/skills/pull/445", Key: "adobe"},
		{Org: "Intel", Repo: "AutoRound", PR: "#2421", URL: "https://github.com/intel/auto-round/pull/2421", Key: "intel"},
		{Org: "IBM", Repo: "Docling Pipelines", PR: "#85", URL: "https://github.com/IBM/docling-pipelines/pull/85", Key: "ibm"},
		{Org: "Uber", Repo: "Kraken", PR: "#716", URL: "https://github.com/uber/kraken/pull/716", Key: "uber"},
		{Org: "Airbnb", Repo: "MagazineLayout", PR: "#169", URL: "https://github.com/airbnb/MagazineLayout/pull/169", Key: "airbnb"},
		{Org: "GitHub", Repo: "GitHub CLI", PR: "#14592", URL: "https://github.com/cli/cli/pull/14592", Key: "github"},
		{Org: "Dell", Repo: "Omnia Containers", PR: "#649", URL: "https://github.com/dell/omnia/pull/649", Key: "dell"},
		{Org: "Salesforce", Repo: "tough-cookie", PR: "#659", URL: "https://github.com/salesforce/tough-cookie/pull/659", Key: "salesforce"},
		{Org: "AMD", Repo: "ROCm HRX System", PR: "#1231", URL: "https://github.com/ROCm/hrx-system/pull/1231", Key: "rocm"},
		{Org: "SAP", Repo: "E-Mobility Simulator", PR: "#2136", URL: "https://github.com/SAP/e-mobility-charging-stations-simulator/pull/2136", Key: "sap"},
		{Org: "Pinterest", Repo: "PyMemcache", PR: "#657", URL: "https://github.com/pinterest/pymemcache/pull/657", Key: "pinterest"},
		{Org: "Qualcomm", Repo: "Commit Msg Check", PR: "#38", URL: "https://github.com/qualcomm/commit-msg-check-action/pull/38", Key: "qualcomm"},
		{Org: "Stripe", Repo: "Stripe React Native", PR: "#2751", URL: "https://github.com/stripe/stripe-react-native/pull/2751", Key: "stripe"},
		{Org: "Docker", Repo: "Docker Docs", PR: "#26266", URL: "https://github.com/docker/docs/pull/26266", Key: "docker"},
		{Org: "Apache", Repo: "HertzBeat", PR: "#4410", URL: "https://github.com/apache/hertzbeat/pull/4410", Key: "apache"},
		{Org: "Cloudflare", Repo: "CFSSL", PR: "#1451", URL: "https://github.com/cloudflare/cfssl/pull/1451", Key: "cloudflare"},
		{Org: "OpenAI", Repo: "Softnet", PR: "#219", URL: "https://github.com/openai/softnet/pull/219", Key: "openai"},
		{Org: "Snapchat", Repo: "ts-inject", PR: "#30", URL: "https://github.com/Snapchat/ts-inject/pull/30", Key: "snapchat"},
		{Org: "Spotify", Repo: "Confidence SDK Swift", PR: "#272", URL: "https://github.com/spotify/confidence-sdk-swift/pull/272", Key: "spotify"},
		{Org: "Cisco", Repo: "cJOSE", PR: "#202", URL: "https://github.com/cisco/cjose/pull/202", Key: "cisco"},
		{Org: "Netflix", Repo: "Atlas", PR: "#2011", URL: "https://github.com/Netflix/atlas/pull/2011", Key: "netflix"},
	}

	experiences := []Experience{
		{
			Company:  "DailyGurus",
			Role:     "Software Developer",
			Period:   "Jul 2026 – Present",
			Location: "Chennai, India",
			Bullets: []string{
				"Engineered scalable full-stack ERP web applications using React.js, Node.js, and TypeScript.",
				"Contributed modular extensions and custom business workflows across database and service layers.",
				"Designed high-performance RESTful APIs handling mission-critical business transactions.",
			},
		},
		{
			Company:  "Data Patterns (India) Ltd.",
			Role:     "Project Intern",
			Period:   "Jan 2026 – May 2026",
			Location: "Chennai, India",
			Bullets: []string{
				"Spearheaded real-time defense communication systems and hardware-in-the-loop telemetry.",
				"Developed high-throughput data processing pipelines for embedded defense electronics.",
				"Optimized data streaming latency by 35% through algorithmic refinement.",
			},
		},
		{
			Company:  "Skill-Lync",
			Role:     "Data Science Intern",
			Period:   "Jun 2024 – Nov 2024",
			Location: "Chennai, India",
			Bullets: []string{
				"Built production machine learning models for customer behavior predictive analytics.",
				"Implemented end-to-end data cleaning, EDA, feature engineering, and cross-validation pipelines.",
				"Automated report generation reducing business analytics cycle time by 40%.",
			},
		},
	}

	hackathons := []Hackathon{
		{
			Title:       "Smart India Hackathon (SIH 2024)",
			Award:       "Grand Finalist · 1st Prize Winner",
			Organizer:   "Ministry of Education & AICTE",
			Year:        "2024",
			Description: "Built AI-Powered Autonomous Crop Health Monitoring and Real-Time Precision Fertilizer Recommendation engine for Indian Agriculture.",
			Tech:        []string{"Python", "Computer Vision", "TensorFlow", "FastAPI", "React"},
		},
		{
			Title:       "Cognizant Digital Disrupt 2024",
			Award:       "National Top 5 Finalist",
			Organizer:   "Cognizant Technology Solutions",
			Year:        "2024",
			Description: "Developed privacy-preserving federated machine learning framework for decentralized clinical record analysis.",
			Tech:        []string{"PyTorch", "Federated Learning", "FastAPI", "Docker"},
		},
		{
			Title:       "TN Police Hackathon 2023",
			Award:       "Special Jury Award for Cybersecurity",
			Organizer:   "Tamil Nadu Police Cyber Crime Division",
			Year:        "2023",
			Description: "Created Tor-Unveil darknet traffic correlation engine identifying rogue relays and Deanonymizing threat actors.",
			Tech:        []string{"Go", "Python", "Wireshark", "Network Forensics"},
		},
	}

	projects := []Project{
		{
			Title:    "301 Super AI Agents Multi-Agent Autonomous Hub",
			Subtitle: "Large-scale autonomous AI agent orchestration mesh with dynamic memory and tool execution",
			Highlights: []string{
				"301 specialized domain agents collaborating via asynchronous DAG pipelines",
				"Sub-150ms semantic memory retrieval across millions of conversation turns",
				"Full self-healing code generation and testing sandbox execution",
			},
			Tags:   []string{"Python", "LangChain", "FastAPI", "VectorDB", "Redis"},
			Github: "https://github.com/reginaldalfret/301-super-ai-agents",
			Demo:   "/projects/301-super-ai-agents",
		},
		{
			Title:    "AI-Powered Crop Disease & Recommendation Engine",
			Subtitle: "Real-time edge computer vision model for automated crop health diagnostics",
			Highlights: []string{
				"97.8% diagnostic accuracy across 42 plant pathology classes",
				"Low-latency on-device inference using quantized TFLite edge models",
				"Awarded Smart India Hackathon Grand Finalist & 1st Prize",
			},
			Tags:   []string{"TensorFlow", "Computer Vision", "Flask", "Tailwind CSS"},
			Github: "https://github.com/reginaldalfret/ai-crop-recommendation",
			Demo:   "/projects/ai-crop-recommendation",
		},
		{
			Title:    "Tor-Unveil Darknet Traffic Correlation & Defense",
			Subtitle: "Cryptographic relay metadata analyzer and forensic de-anonymization platform",
			Highlights: []string{
				"Statistical packet flow correlation for identifying rogue Tor circuits",
				"Interactive topological graph visualizer of encrypted node hops",
				"Presented at TN Cyber Crime Department Technical Forum",
			},
			Tags:   []string{"Go", "Network Security", "React", "D3.js"},
			Github: "https://github.com/reginaldalfret/tor-unveil",
			Demo:   "/projects/tor-unveil",
		},
	}

	skills := []SkillCategory{
		{
			Category: "Languages & Runtimes",
			Skills:   []string{"Go (Golang)", "Python", "TypeScript", "JavaScript", "C#", "C++", "Java", "Ruby", "PHP", "SQL"},
		},
		{
			Category: "Frameworks & Engines",
			Skills:   []string{"Go Buffalo", "Next.js", "Nuxt 3", "Django", "Spring Boot", "Laravel", "Ruby on Rails", "Angular", "Blazor WASM", "FastAPI", "Express.js"},
		},
		{
			Category: "AI / Data Science & ML",
			Skills:   []string{"PyTorch", "TensorFlow", "LangChain", "Hugging Face", "Pandas", "NumPy", "OpenCV", "Scikit-Learn", "Vector Databases"},
		},
		{
			Category: "DevOps, Cloud & Architecture",
			Skills:   []string{"Docker", "Kubernetes", "AWS", "Google Cloud", "Azure", "PostgreSQL", "Redis", "Kafka", "Git & CI/CD"},
		},
	}

	achievements := []Achievement{
		{Title: "30 Merged PRs Across Big Tech & Open-Source Giants", Description: "Contributions accepted by Apple, Google, Microsoft, Meta, AWS, NVIDIA, NASA, Netflix, Linux Foundation & Apache."},
		{Title: "Smart India Hackathon 2024 Winner", Description: "1st Prize in India's biggest national hackathon organized by AICTE & Ministry of Education."},
		{Title: "Published Research & Technical Contributions", Description: "Authored papers on real-time neural edge processing and cryptographic forensic correlation."},
	}

	certifications := []Certification{
		{Name: "Google Cloud Certified Professional Data Engineer", Issuer: "Google Cloud", Date: "2024"},
		{Name: "AWS Certified Machine Learning – Specialty", Issuer: "Amazon Web Services", Date: "2024"},
		{Name: "Deep Learning Specialization", Issuer: "DeepLearning.AI", Date: "2023"},
		{Name: "Certified Kubernetes Application Developer (CKAD)", Issuer: "Cloud Native Computing Foundation", Date: "2024"},
	}

	stacks := []StackItem{
		{
			Name:        "Next.js 14",
			Subtitle:    "React 18 · App Router · TypeScript · Tailwind CSS · Vercel Edge.",
			Route:       "/",
			Repo:        "https://github.com/reginaldalfret/portfolio-website",
			Badge:       "ACTIVE CORE",
			BadgeClass:  "bg-emerald-500/20 text-emerald-300 border border-emerald-500/30",
			BorderClass: "border-cyan-500/30",
			Icon:        "N",
			IsCurrent:   false,
		},
		{
			Name:        "Nuxt 3",
			Subtitle:    "Vue 3 · Composition API · Nitro Engine · Vite · SSR.",
			Route:       "/stacks/nuxt",
			Repo:        "https://github.com/reginaldalfret/portfolio-nuxt",
			Badge:       "VUE 3",
			BadgeClass:  "bg-emerald-500/10 text-emerald-400 border border-emerald-500/20",
			BorderClass: "hover:border-emerald-500/40",
			Icon:        "▲",
			IsCurrent:   false,
		},
		{
			Name:        "Django",
			Subtitle:    "Python 3.13 · Django SSR Templates · Tailwind CSS · Static Engine.",
			Route:       "/stacks/django",
			Repo:        "https://github.com/reginaldalfret/portfolio-django",
			Badge:       "PYTHON 3.13",
			BadgeClass:  "bg-amber-500/10 text-amber-400 border border-amber-500/20",
			BorderClass: "hover:border-amber-500/40",
			Icon:        "DJ",
			IsCurrent:   false,
		},
		{
			Name:        "Spring Boot",
			Subtitle:    "Java 21 LTS · Thymeleaf SSR · Spring MVC · Maven · Tomcat.",
			Route:       "/stacks/springboot",
			Repo:        "https://github.com/reginaldalfret/portfolio-springboot",
			Badge:       "JAVA 21",
			BadgeClass:  "bg-lime-500/10 text-lime-400 border border-lime-500/20",
			BorderClass: "hover:border-lime-500/40",
			Icon:        "🍃",
			IsCurrent:   false,
		},
		{
			Name:        "Laravel",
			Subtitle:    "PHP 8.4 · Blade Templates · Artisan · Vite · Tailwind CSS.",
			Route:       "/stacks/laravel",
			Repo:        "https://github.com/reginaldalfret/portfolio-laravel",
			Badge:       "PHP 8.4",
			BadgeClass:  "bg-red-500/10 text-red-400 border border-red-500/20",
			BorderClass: "hover:border-red-500/40",
			Icon:        "L",
			IsCurrent:   false,
		},
		{
			Name:        "Ruby on Rails",
			Subtitle:    "Ruby 3.4 · Rails 8 · ActionView ERB SSR · Puma Server.",
			Route:       "/stacks/rails",
			Repo:        "https://github.com/reginaldalfret/portfolio-rails",
			Badge:       "RAILS 8",
			BadgeClass:  "bg-rose-500/10 text-rose-400 border border-rose-500/20",
			BorderClass: "hover:border-rose-500/40",
			Icon:        "💎",
			IsCurrent:   false,
		},
		{
			Name:        "Angular",
			Subtitle:    "Angular 19/22 · Signals · Standalone Components · Tailwind CSS.",
			Route:       "/stacks/angular",
			Repo:        "https://github.com/reginaldalfret/portfolio-angular",
			Badge:       "TS 6.0",
			BadgeClass:  "bg-red-500/10 text-red-400 border border-red-500/20",
			BorderClass: "hover:border-red-500/40",
			Icon:        "A",
			IsCurrent:   false,
		},
		{
			Name:        "Blazor WASM",
			Subtitle:    ".NET 8 · C# Client-Side WASM · Standalone SPA · Tailwind CSS.",
			Route:       "/stacks/blazor",
			Repo:        "https://github.com/reginaldalfret/portfolio-blazor",
			Badge:       ".NET 8",
			BadgeClass:  "bg-purple-500/10 text-purple-400 border border-purple-500/20",
			BorderClass: "hover:border-purple-500/40",
			Icon:        "B",
			IsCurrent:   false,
		},
		{
			Name:        "Go Buffalo",
			Subtitle:    "Go 1.27 · Buffalo Web Framework · Plush SSR · Gorilla Mux.",
			Route:       "/stacks/buffalo",
			Repo:        "https://github.com/reginaldalfret/portfolio-buffalo",
			Badge:       "GO 1.27",
			BadgeClass:  "bg-cyan-500/20 text-cyan-300 border border-cyan-500/40",
			BorderClass: "border-cyan-500/40",
			Icon:        "🦬",
			IsCurrent:   true,
		},
	}

	return map[string]interface{}{
		"contributions":  contributions,
		"experiences":    experiences,
		"hackathons":      hackathons,
		"projects":       projects,
		"skills":         skills,
		"achievements":   achievements,
		"certifications": certifications,
		"stacks":         stacks,
	}
}
