package main

type Country struct {
	Name string
	Code string
	Flag string
}

var countries = map[string]Country{
	"AFG": {
		Name: "Afghanistan",
		Code: "AFG",
		Flag: ":flag-af:",
	},
	"ALB": {
		Name: "Albania",
		Code: "ALB",
		Flag: ":flag-al:",
	},
	"ALG": {
		Name: "Algeria",
		Code: "ALG",
		Flag: ":flag-dz:",
	},
	"ASA": {
		Name: "American Samoa",
		Code: "ASA",
		Flag: ":flag-as:",
	},
	"AND": {
		Name: "Andorra",
		Code: "AND",
		Flag: ":flag-ad:",
	},
	"ANG": {
		Name: "Angola",
		Code: "ANG",
		Flag: ":flag-ao:",
	},
	"AIA": {
		Name: "Anguilla",
		Code: "AIA",
		Flag: ":flag-ai:",
	},
	"ATG": {
		Name: "Antigua and Barbuda",
		Code: "ATG",
		Flag: ":flag-ag:",
	},
	"ARG": {
		Name: "Argentina",
		Code: "ARG",
		Flag: ":flag-ar:",
	},
	"ARM": {
		Name: "Armenia",
		Code: "ARM",
		Flag: ":flag-am:",
	},
	"ARU": {
		Name: "Aruba",
		Code: "ARU",
		Flag: ":flag-aw:",
	},
	"AUS": {
		Name: "Australia",
		Code: "AUS",
		Flag: ":flag-au:",
	},
	"AUT": {
		Name: "Austria",
		Code: "AUT",
		Flag: ":flag-at:",
	},
	"AZE": {
		Name: "Azerbaijan",
		Code: "AZE",
		Flag: ":flag-az:",
	},
	"BAH": {
		Name: "The Bahamas",
		Code: "BAH",
		Flag: ":flag-bs:",
	},
	"BHR": {
		Name: "Bahrain",
		Code: "BHR",
		Flag: ":flag-bh:",
	},
	"BAN": {
		Name: "Bangladesh",
		Code: "BAN",
		Flag: ":flag-bd:",
	},
	"BRB": {
		Name: "Barbados",
		Code: "BRB",
		Flag: ":flag-bb:",
	},
	"BLR": {
		Name: "Belarus",
		Code: "BLR",
		Flag: ":flag-by:",
	},
	"BEL": {
		Name: "Belgium",
		Code: "BEL",
		Flag: ":flag-be:",
	},
	"BLZ": {
		Name: "Belize",
		Code: "BLZ",
		Flag: ":flag-bz:",
	},
	"BEN": {
		Name: "Benin",
		Code: "BEN",
		Flag: ":flag-bj:",
	},
	"BER": {
		Name: "Bermuda",
		Code: "BER",
		Flag: ":flag-bm:",
	},
	"BHU": {
		Name: "Bhutan",
		Code: "BHU",
		Flag: ":flag-bt:",
	},
	"BOL": {
		Name: "Bolivia",
		Code: "BOL",
		Flag: ":flag-bo:",
	},
	"BIH": {
		Name: "Bosnia and Herzegovina",
		Code: "BIH",
		Flag: ":flag-ba:",
	},
	"BOT": {
		Name: "Botswana",
		Code: "BOT",
		Flag: ":flag-bw:",
	},
	"BRA": {
		Name: "Brazil",
		Code: "BRA",
		Flag: ":flag-br:",
	},
	"VGB": {
		Name: "British Virgin Islands",
		Code: "VGB",
		Flag: ":flag-vg:",
	},
	"BRU": {
		Name: "Brunei",
		Code: "BRU",
		Flag: ":flag-bn:",
	},
	"BUL": {
		Name: "Bulgaria",
		Code: "BUL",
		Flag: ":flag-bg:",
	},
	"BFA": {
		Name: "Burkina Faso",
		Code: "BFA",
		Flag: ":flag-bf:",
	},
	"BDI": {
		Name: "Burundi",
		Code: "BDI",
		Flag: ":flag-bi:",
	},
	"CAM": {
		Name: "Cambodia",
		Code: "CAM",
		Flag: ":flag-kh:",
	},
	"CMR": {
		Name: "Cameroon",
		Code: "CMR",
		Flag: ":flag-cm:",
	},
	"CAN": {
		Name: "Canada",
		Code: "CAN",
		Flag: ":flag-ca:",
	},
	"CPV": {
		Name: "Cape Verde",
		Code: "CPV",
		Flag: ":flag-cv:",
	},
	"CAY": {
		Name: "Cayman Islands",
		Code: "CAY",
		Flag: ":flag-ky:",
	},
	"CTA": {
		Name: "Central African Republic",
		Code: "CTA",
		Flag: ":flag-cf:",
	},
	"CHA": {
		Name: "Chad",
		Code: "CHA",
		Flag: ":flag-td:",
	},
	"CHI": {
		Name: "Chile",
		Code: "CHI",
		Flag: ":flag-cl:",
	},
	"CHN": {
		Name: "China",
		Code: "CHN",
		Flag: ":flag-cn:",
	},
	"COL": {
		Name: "Colombia",
		Code: "COL",
		Flag: ":flag-co:",
	},
	"COM": {
		Name: "Comoros",
		Code: "COM",
		Flag: ":flag-km:",
	},
	"COD": {
		Name: "Democratic Republic of the Congo",
		Code: "COD",
		Flag: ":flag-cd:",
	},
	"CGO": {
		Name: "Republic of the Congo",
		Code: "CGO",
		Flag: ":flag-cg:",
	},
	"COK": {
		Name: "Cook Islands",
		Code: "COK",
		Flag: ":flag-ck:",
	},
	"CRC": {
		Name: "Costa Rica",
		Code: "CRC",
		Flag: ":flag-cr:",
	},
	"CIV": {
		Name: "Ivory Coast",
		Code: "CIV",
		Flag: ":flag-ci:",
	},
	"CRO": {
		Name: "Croatia",
		Code: "CRO",
		Flag: ":flag-hr:",
	},
	"CUB": {
		Name: "Cuba",
		Code: "CUB",
		Flag: ":flag-cu:",
	},
	"CYP": {
		Name: "Cyprus",
		Code: "CYP",
		Flag: ":flag-cy:",
	},
	"CZE": {
		Name: "Czech Republic",
		Code: "CZE",
		Flag: ":flag-cz:",
	},
	"DEN": {
		Name: "Denmark",
		Code: "DEN",
		Flag: ":flag-dk:",
	},
	"DJI": {
		Name: "Djibouti",
		Code: "DJI",
		Flag: ":flag-dj:",
	},
	"DMA": {
		Name: "Dominica",
		Code: "DMA",
		Flag: ":flag-dm:",
	},
	"DOM": {
		Name: "Dominican Republic",
		Code: "DOM",
		Flag: ":flag-do:",
	},
	"ECU": {
		Name: "Ecuador",
		Code: "ECU",
		Flag: ":flag-ec:",
	},
	"EGY": {
		Name: "Egypt",
		Code: "EGY",
		Flag: ":flag-eg:",
	},
	"SLV": {
		Name: "El Salvador",
		Code: "SLV",
		Flag: ":flag-sv:",
	},
	"ENG": {
		Name: "England",
		Code: "ENG",
		Flag: ":flag-england:",
	},
	"EQG": {
		Name: "Equatorial Guinea",
		Code: "EQG",
		Flag: ":flag-gq:",
	},
	"ERI": {
		Name: "Eritrea",
		Code: "ERI",
		Flag: ":flag-er:",
	},
	"EST": {
		Name: "Estonia",
		Code: "EST",
		Flag: ":flag-ee:",
	},
	"ETH": {
		Name: "Ethiopia",
		Code: "ETH",
		Flag: ":flag-et:",
	},
	"FRO": {
		Name: "Faroe Islands",
		Code: "FRO",
		Flag: ":flag-fo:",
	},
	"FIJ": {
		Name: "Fiji",
		Code: "FIJ",
		Flag: ":flag-fj:",
	},
	"FIN": {
		Name: "Finland",
		Code: "FIN",
		Flag: ":flag-fi:",
	},
	"FRA": {
		Name: "France",
		Code: "FRA",
		Flag: ":flag-fr:",
	},
	"TAH": {
		Name: "French Polynesia",
		Code: "TAH",
		Flag: ":flag-pf:",
	},
	"GAB": {
		Name: "Gabon",
		Code: "GAB",
		Flag: ":flag-ga:",
	},
	"GAM": {
		Name: "The Gambia",
		Code: "GAM",
		Flag: ":flag-gm:",
	},
	"GEO": {
		Name: "Georgia (country)",
		Code: "GEO",
		Flag: ":flag-ge:",
	},
	"GER": {
		Name: "Germany",
		Code: "GER",
		Flag: ":flag-de:",
	},
	"GHA": {
		Name: "Ghana",
		Code: "GHA",
		Flag: ":flag-gh:",
	},
	"GIB": {
		Name: "Gibraltar",
		Code: "GIB",
		Flag: ":flag-gi:",
	},
	"GRE": {
		Name: "Greece",
		Code: "GRE",
		Flag: ":flag-gr:",
	},
	"GRN": {
		Name: "Grenada",
		Code: "GRN",
		Flag: ":flag-gd:",
	},
	"GUM": {
		Name: "Guam",
		Code: "GUM",
		Flag: ":flag-gu:",
	},
	"GUA": {
		Name: "Guatemala",
		Code: "GUA",
		Flag: ":flag-gt:",
	},
	"GUI": {
		Name: "Guinea",
		Code: "GUI",
		Flag: ":flag-gn:",
	},
	"GNB": {
		Name: "Guinea-Bissau",
		Code: "GNB",
		Flag: ":flag-gw:",
	},
	"GUY": {
		Name: "Guyana",
		Code: "GUY",
		Flag: ":flag-gy:",
	},
	"HAI": {
		Name: "Haiti",
		Code: "HAI",
		Flag: ":flag-ht:",
	},
	"HON": {
		Name: "Honduras",
		Code: "HON",
		Flag: ":flag-hn:",
	},
	"HKG": {
		Name: "Hong Kong",
		Code: "HKG",
		Flag: ":flag-hk:",
	},
	"HUN": {
		Name: "Hungary",
		Code: "HUN",
		Flag: ":flag-hu:",
	},
	"ISL": {
		Name: "Iceland",
		Code: "ISL",
		Flag: ":flag-is:",
	},
	"IND": {
		Name: "India",
		Code: "IND",
		Flag: ":flag-in:",
	},
	"IDN": {
		Name: "Indonesia",
		Code: "IDN",
		Flag: ":flag-id:",
	},
	"IRN": {
		Name: "Iran",
		Code: "IRN",
		Flag: ":flag-ir:",
	},
	"IRQ": {
		Name: "Iraq",
		Code: "IRQ",
		Flag: ":flag-iq:",
	},
	"IRL": {
		Name: "Republic of Ireland",
		Code: "IRL",
		Flag: ":flag-ie:",
	},
	"ISR": {
		Name: "Israel",
		Code: "ISR",
		Flag: ":flag-il:",
	},
	"ITA": {
		Name: "Italy",
		Code: "ITA",
		Flag: ":flag-it:",
	},
	"JAM": {
		Name: "Jamaica",
		Code: "JAM",
		Flag: ":flag-jm:",
	},
	"JPN": {
		Name: "Japan",
		Code: "JPN",
		Flag: ":flag-jp:",
	},
	"JOR": {
		Name: "Jordan",
		Code: "JOR",
		Flag: ":flag-jo:",
	},
	"KAZ": {
		Name: "Kazakhstan",
		Code: "KAZ",
		Flag: ":flag-kz:",
	},
	"KEN": {
		Name: "Kenya",
		Code: "KEN",
		Flag: ":flag-ke:",
	},
	"PRK": {
		Name: "North Korea",
		Code: "PRK",
		Flag: ":flag-kp:",
	},
	"KOR": {
		Name: "South Korea",
		Code: "KOR",
		Flag: ":flag-kr:",
	},
	"KUW": {
		Name: "Kuwait",
		Code: "KUW",
		Flag: ":flag-kw:",
	},
	"KGZ": {
		Name: "Kyrgyzstan",
		Code: "KGZ",
		Flag: ":flag-kg:",
	},
	"LAO": {
		Name: "Laos",
		Code: "LAO",
		Flag: ":flag-la:",
	},
	"LVA": {
		Name: "Latvia",
		Code: "LVA",
		Flag: ":flag-lv:",
	},
	"LBN": {
		Name: "Lebanon",
		Code: "LBN",
		Flag: ":flag-lb:",
	},
	"LES": {
		Name: "Lesotho",
		Code: "LES",
		Flag: ":flag-ls:",
	},
	"LBR": {
		Name: "Liberia",
		Code: "LBR",
		Flag: ":flag-lr:",
	},
	"LBY": {
		Name: "Libya",
		Code: "LBY",
		Flag: ":flag-ly:",
	},
	"LIE": {
		Name: "Liechtenstein",
		Code: "LIE",
		Flag: ":flag-li:",
	},
	"LTU": {
		Name: "Lithuania",
		Code: "LTU",
		Flag: ":flag-lt:",
	},
	"LUX": {
		Name: "Luxembourg",
		Code: "LUX",
		Flag: ":flag-lu:",
	},
	"MAC": {
		Name: "Macau",
		Code: "MAC",
		Flag: ":flag-mo:",
	},
	"MKD": {
		Name: "Republic of Macedonia",
		Code: "MKD",
		Flag: ":flag-mk:",
	},
	"MAD": {
		Name: "Madagascar",
		Code: "MAD",
		Flag: ":flag-mg:",
	},
	"MWI": {
		Name: "Malawi",
		Code: "MWI",
		Flag: ":flag-mw:",
	},
	"MAS": {
		Name: "Malaysia",
		Code: "MAS",
		Flag: ":flag-my:",
	},
	"MDV": {
		Name: "Maldives",
		Code: "MDV",
		Flag: ":flag-mv:",
	},
	"MLI": {
		Name: "Mali",
		Code: "MLI",
		Flag: ":flag-ml:",
	},
	"MLT": {
		Name: "Malta",
		Code: "MLT",
		Flag: ":flag-mt:",
	},
	"MTN": {
		Name: "Mauritania",
		Code: "MTN",
		Flag: ":flag-mr:",
	},
	"MRI": {
		Name: "Mauritius",
		Code: "MRI",
		Flag: ":flag-mu:",
	},
	"MEX": {
		Name: "Mexico",
		Code: "MEX",
		Flag: ":flag-mx:",
	},
	"MDA": {
		Name: "Moldova",
		Code: "MDA",
		Flag: ":flag-md:",
	},
	"MNG": {
		Name: "Mongolia",
		Code: "MNG",
		Flag: ":flag-mn:",
	},
	"MNE": {
		Name: "Montenegro",
		Code: "MNE",
		Flag: ":flag-me:",
	},
	"MSR": {
		Name: "Montserrat",
		Code: "MSR",
		Flag: ":flag-ms:",
	},
	"MAR": {
		Name: "Morocco",
		Code: "MAR",
		Flag: ":flag-ma:",
	},
	"MOZ": {
		Name: "Mozambique",
		Code: "MOZ",
		Flag: ":flag-mz:",
	},
	"MYA": {
		Name: "Myanmar",
		Code: "MYA",
		Flag: ":flag-mm:",
	},
	"NAM": {
		Name: "Namibia",
		Code: "NAM",
		Flag: ":flag-na:",
	},
	"NEP": {
		Name: "Nepal",
		Code: "NEP",
		Flag: ":flag-np:",
	},
	"NED": {
		Name: "Netherlands",
		Code: "NED",
		Flag: ":flag-nl:",
	},
	"NCL": {
		Name: "New Caledonia",
		Code: "NCL",
		Flag: ":flag-nc:",
	},
	"NZL": {
		Name: "New Zealand",
		Code: "NZL",
		Flag: ":flag-nz:",
	},
	"NCA": {
		Name: "Nicaragua",
		Code: "NCA",
		Flag: ":flag-ni:",
	},
	"NIG": {
		Name: "Niger",
		Code: "NIG",
		Flag: ":flag-ne:",
	},
	"NGA": {
		Name: "Nigeria",
		Code: "NGA",
		Flag: ":flag-ng:",
	},
	"NOR": {
		Name: "Norway",
		Code: "NOR",
		Flag: ":flag-no:",
	},
	"OMA": {
		Name: "Oman",
		Code: "OMA",
		Flag: ":flag-om:",
	},
	"PAK": {
		Name: "Pakistan",
		Code: "PAK",
		Flag: ":flag-pk:",
	},
	"PLE": {
		Name: "Palestinian National Authority",
		Code: "PLE",
		Flag: ":flag-ps:",
	},
	"PAN": {
		Name: "Panama",
		Code: "PAN",
		Flag: ":flag-pa:",
	},
	"PNG": {
		Name: "Papua New Guinea",
		Code: "PNG",
		Flag: ":flag-pg:",
	},
	"PAR": {
		Name: "Paraguay",
		Code: "PAR",
		Flag: ":flag-py:",
	},
	"PER": {
		Name: "Peru",
		Code: "PER",
		Flag: ":flag-pe:",
	},
	"PHI": {
		Name: "Philippines",
		Code: "PHI",
		Flag: ":flag-ph:",
	},
	"POL": {
		Name: "Poland",
		Code: "POL",
		Flag: ":flag-pl:",
	},
	"POR": {
		Name: "Portugal",
		Code: "POR",
		Flag: ":flag-pt:",
	},
	"PUR": {
		Name: "Puerto Rico",
		Code: "PUR",
		Flag: ":flag-pr:",
	},
	"QAT": {
		Name: "Qatar",
		Code: "QAT",
		Flag: ":flag-qa:",
	},
	"ROU": {
		Name: "Romania",
		Code: "ROU",
		Flag: ":flag-ro:",
	},
	"RUS": {
		Name: "Russia",
		Code: "RUS",
		Flag: ":flag-ru:",
	},
	"RWA": {
		Name: "Rwanda",
		Code: "RWA",
		Flag: ":flag-rw:",
	},
	"SKN": {
		Name: "Saint Kitts and Nevis",
		Code: "SKN",
		Flag: ":flag-kn:",
	},
	"LCA": {
		Name: "Saint Lucia",
		Code: "LCA",
		Flag: ":flag-lc:",
	},
	"VIN": {
		Name: "Saint Vincent and the Grenadines",
		Code: "VIN",
		Flag: ":flag-vc:",
	},
	"SAM": {
		Name: "Samoa",
		Code: "SAM",
		Flag: ":flag-ws:",
	},
	"SMR": {
		Name: "San Marino",
		Code: "SMR",
		Flag: ":flag-sm:",
	},
	"STP": {
		Name: "São Tomé and Príncipe",
		Code: "STP",
		Flag: ":flag-st:",
	},
	"KSA": {
		Name: "Saudi Arabia",
		Code: "KSA",
		Flag: ":flag-sa:",
	},
	"SCO": {
		Name: "Scotland",
		Code: "SCO",
		Flag: ":flag-scotland:",
	},
	"SEN": {
		Name: "Senegal",
		Code: "SEN",
		Flag: ":flag-sn:",
	},
	"SRB": {
		Name: "Serbia",
		Code: "SRB",
		Flag: ":flag-rs:",
	},
	"SEY": {
		Name: "Seychelles",
		Code: "SEY",
		Flag: ":flag-sc:",
	},
	"SLE": {
		Name: "Sierra Leone",
		Code: "SLE",
		Flag: ":flag-sl:",
	},
	"SIN": {
		Name: "Singapore",
		Code: "SIN",
		Flag: ":flag-sg:",
	},
	"SVK": {
		Name: "Slovakia",
		Code: "SVK",
		Flag: ":flag-sk:",
	},
	"SVN": {
		Name: "Slovenia",
		Code: "SVN",
		Flag: ":flag-si:",
	},
	"SOL": {
		Name: "Solomon Islands",
		Code: "SOL",
		Flag: ":flag-sb:",
	},
	"SOM": {
		Name: "Somalia",
		Code: "SOM",
		Flag: ":flag-so:",
	},
	"RSA": {
		Name: "South Africa",
		Code: "RSA",
		Flag: ":flag-za:",
	},
	"ESP": {
		Name: "Spain",
		Code: "ESP",
		Flag: ":flag-es:",
	},
	"SRI": {
		Name: "Sri Lanka",
		Code: "SRI",
		Flag: ":flag-lk:",
	},
	"SDN": {
		Name: "Sudan",
		Code: "SDN",
		Flag: ":flag-sd:",
	},
	"SUR": {
		Name: "Suriname",
		Code: "SUR",
		Flag: ":flag-sr:",
	},
	"SWZ": {
		Name: "Eswatini",
		Code: "SWZ",
		Flag: ":flag-sz:",
	},
	"SWE": {
		Name: "Sweden",
		Code: "SWE",
		Flag: ":flag-se:",
	},
	"SUI": {
		Name: "Switzerland",
		Code: "SUI",
		Flag: ":flag-ch:",
	},
	"SYR": {
		Name: "Syria",
		Code: "SYR",
		Flag: ":flag-sy:",
	},
	"TPE": {
		Name: "Taiwan",
		Code: "TPE",
		Flag: ":flag-tw:",
	},
	"TJK": {
		Name: "Tajikistan",
		Code: "TJK",
		Flag: ":flag-tj:",
	},
	"TAN": {
		Name: "Tanzania",
		Code: "TAN",
		Flag: ":flag-tz:",
	},
	"THA": {
		Name: "Thailand",
		Code: "THA",
		Flag: ":flag-th:",
	},
	"TLS": {
		Name: "East Timor",
		Code: "TLS",
		Flag: ":flag-tl:",
	},
	"TOG": {
		Name: "Togo",
		Code: "TOG",
		Flag: ":flag-tg:",
	},
	"TGA": {
		Name: "Tonga",
		Code: "TGA",
		Flag: ":flag-to:",
	},
	"TRI": {
		Name: "Trinidad and Tobago",
		Code: "TRI",
		Flag: ":flag-tt:",
	},
	"TUN": {
		Name: "Tunisia",
		Code: "TUN",
		Flag: ":flag-tn:",
	},
	"TUR": {
		Name: "Turkey",
		Code: "TUR",
		Flag: ":flag-tr:",
	},
	"TKM": {
		Name: "Turkmenistan",
		Code: "TKM",
		Flag: ":flag-tm:",
	},
	"TCA": {
		Name: "Turks and Caicos Islands",
		Code: "TCA",
		Flag: ":flag-tc:",
	},
	"UGA": {
		Name: "Uganda",
		Code: "UGA",
		Flag: ":flag-ug:",
	},
	"UKR": {
		Name: "Ukraine",
		Code: "UKR",
		Flag: ":flag-ua:",
	},
	"UAE": {
		Name: "United Arab Emirates",
		Code: "UAE",
		Flag: ":flag-ae:",
	},
	"USA": {
		Name: "United States",
		Code: "USA",
		Flag: ":flag-us:",
	},
	"VIR": {
		Name: "United States Virgin Islands",
		Code: "VIR",
		Flag: ":flag-vi:",
	},
	"URU": {
		Name: "Uruguay",
		Code: "URU",
		Flag: ":flag-uy:",
	},
	"UZB": {
		Name: "Uzbekistan",
		Code: "UZB",
		Flag: ":flag-uz:",
	},
	"VAN": {
		Name: "Vanuatu",
		Code: "VAN",
		Flag: ":flag-vu:",
	},
	"VEN": {
		Name: "Venezuela",
		Code: "VEN",
		Flag: ":flag-ve:",
	},
	"VIE": {
		Name: "Vietnam",
		Code: "VIE",
		Flag: ":flag-vn:",
	},
	"WAL": {
		Name: "Wales",
		Code: "WAL",
		Flag: ":flag-wales:",
	},
	"ESH": {
		Name: "Western Sahara",
		Code: "ESH",
		Flag: ":flag-eh:",
	},
	"YEM": {
		Name: "Yemen",
		Code: "YEM",
		Flag: ":flag-ye:",
	},
	"ZAM": {
		Name: "Zambia",
		Code: "ZAM",
		Flag: ":flag-zm:",
	},
	"ZIM": {
		Name: "Zimbabwe",
		Code: "ZIM",
		Flag: ":flag-zw:",
	},
}
