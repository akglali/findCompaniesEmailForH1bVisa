package googleMapsSearch

var cityAndState = map[string][]string{
	"Alabama": {
		"Birmingham",
		"Montgomery",
		"Mobile",
		"Huntsville",
		"Tuscaloosa",
		"Auburn",
		"Decatur",
		"Dothan",
		"Gadsden",
		"Hoover",
	},
	"Alaska": {
		"Anchorage",
		"Fairbanks",
		"Juneau",
		"Wasilla",
		"Palmer",
		"Kenai",
		"Soldotna",
		"Sitka",
		"Kodiak",
		"Eagle River",
	},
	"Arizona": {
		"Phoenix",
		"Tucson",
		"Mesa",
		"Chandler",
		"Glendale",
		"Scottsdale",
		"Gilbert",
		"Tempe",
		"Peoria",
		"Surprise",
	},
	"Arkansas": {
		"Little Rock",
		"Fort Smith",
		"Fayetteville",
		"Jonesboro",
		"North Little Rock",
		"Rogers",
		"Pine Bluff",
		"Bentonville",
		"Hot Springs",
		"Springdale",
	},
	"California": {
		"Los Angeles",
		"San Diego",
		"San Jose",
		"San Francisco",
		"Fresno",
		"Sacramento",
		"Long Beach",
		"Oakland",
		"Bakersfield",
		"Anaheim",
	},
	"Colorado": {
		"Denver",
		"Colorado Springs",
		"Aurora",
		"Fort Collins",
		"Lakewood",
		"Thornton",
		"Arvada",
		"Westminster",
		"Pueblo",
		"Centennial",
	},
	"Connecticut": {
		"Bridgeport",
		"New Haven",
		"Hartford",
		"Stamford",
		"Waterbury",
		"Norwalk",
		"Danbury",
		"New Britain",
		"Bristol",
		"Meriden",
	},
	"Delaware": {
		"Wilmington",
		"Dover",
		"Newark",
		"Middletown",
		"Smyrna",
		"Milford",
		"Seaford",
		"Elsmere",
		"Georgetown",
		"New Castle",
	},
	"Florida": {
		"Jacksonville",
		"Miami",
		"Tampa",
		"Orlando",
		"St. Petersburg",
		"Hialeah",
		"Tallahassee",
		"Fort Lauderdale",
		"Port St. Lucie",
		"Pembroke Pines",
	},
	"Georgia": {
		"Atlanta",
		"Augusta",
		"Columbus",
		"Savannah",
		"Athens",
		"Macon",
		"Sandy Springs",
		"Roswell",
		"Johns Creek",
		"Albany",
	},
	"Hawaii": {
		"Honolulu",
		"East Honolulu",
		"Pearl City",
		"Hilo",
		"Kailua",
		"Waipahu",
		"Kaneohe",
		"Mililani Town",
		"Kahului",
		"Kihei",
	},
	"Idaho": {
		"Boise",
		"Nampa",
		"Meridian",
		"Idaho Falls",
		"Pocatello",
		"Caldwell",
		"Coeur d'Alene",
		"Twin Falls",
		"Lewiston",
		"Post Falls",
	},
	"Illinois": {
		"Chicago",
		"Aurora",
		"Rockford",
		"Joliet",
		"Naperville",
		"Springfield",
		"Elgin",
		"Peoria",
		"Waukegan",
		"Cicero",
	},
	"Indiana": {
		"Indianapolis",
		"Fort Wayne",
		"Evansville",
		"South Bend",
		"Carmel",
		"Fishers",
		"Bloomington",
		"Hammond",
		"Gary",
		"Lafayette",
	},
	"Iowa": {
		"Des Moines",
		"Cedar Rapids",
		"Davenport",
		"Sioux City",
		"Waterloo",
		"Iowa City",
		"Council Bluffs",
		"Ames",
		"West Des Moines",
		"Dubuque",
	},
	"Kansas": {
		"Wichita",
		"Overland Park",
		"Kansas City",
		"Topeka",
		"Olathe",
		"Lawrence",
		"Shawnee",
		"Manhattan",
		"Lenexa",
		"Salina",
	},
	"Kentucky": {
		"Louisville",
		"Lexington",
		"Bowling Green",
		"Owensboro",
		"Covington",
		"Hopkinsville",
		"Richmond",
		"Florence",
		"Georgetown",
		"Henderson",
	},
	"Louisiana": {
		"New Orleans",
		"Baton Rouge",
		"Shreveport",
		"Lafayette",
		"Lake Charles",
		"Kenner",
		"Bossier City",
		"Monroe",
		"Alexandria",
		"New Iberia",
	},
	"Maine": {
		"Portland",
		"Lewiston",
		"Bangor",
		"South Portland",
		"Auburn",
		"Biddeford",
		"Sanford",
		"Westbrook",
		"Augusta",
		"Waterville",
	},
	"Maryland": {
		"Baltimore",
		"Frederick",
		"Rockville",
		"Gaithersburg",
		"Bowie",
		"Hagerstown",
		"Annapolis",
		"College Park",
		"Laurel",
		"Greenbelt",
	},
	"Massachusetts": {
		"Boston",
		"Worcester",
		"Springfield",
		"Lowell",
		"Cambridge",
		"New Bedford",
		"Brockton",
		"Quincy",
		"Lynn",
		"Newton",
	},
	"Michigan": {
		"Detroit",
		"Grand Rapids",
		"Warren",
		"Sterling Heights",
		"Ann Arbor",
		"Lansing",
		"Flint",
		"Dearborn",
		"Livonia",
		"Westland",
	},
	"Minnesota": {
		"Minneapolis",
		"St. Paul",
		"Rochester",
		"Duluth",
		"Bloomington",
		"Brooklyn Park",
		"Plymouth",
		"Eagan",
		"Coon Rapids",
		"Burnsville",
	},
	"Mississippi": {
		"Jackson",
		"Gulfport",
		"Southaven",
		"Hattiesburg",
		"Biloxi",
		"Meridian",
		"Greenville",
		"Tupelo",
		"Olive Branch",
		"Horn Lake",
	},
	"Missouri": {
		"Kansas City",
		"St. Louis",
		"Springfield",
		"Independence",
		"Columbia",
		"Lee's Summit",
		"O'Fallon",
		"St. Joseph",
		"St. Charles",
		"Blue Springs",
	},
	"Montana": {
		"Billings",
		"Missoula",
		"Great Falls",
		"Bozeman",
		"Butte",
		"Helena",
		"Kalispell",
		"Havre",
		"Anaconda",
		"Miles City",
	},
	"Nebraska": {
		"Omaha",
		"Lincoln",
		"Bellevue",
		"Grand Island",
		"Kearney",
		"Fremont",
		"Hastings",
		"North Platte",
		"Norfolk",
		"Columbus",
	},
	"Nevada": {
		"Las Vegas",
		"Reno",
		"Henderson",
		"North Las Vegas",
		"Sparks",
		"Carson City",
		"Elko",
		"Fernley",
		"Wells",
		"Winnemucca",
	},
	"New Hampshire": {
		"Manchester",
		"Nashua",
		"Concord",
		"Dover",
		"Rochester",
		"Keene",
		"Salem",
		"Portsmouth",
		"Londonderry",
		"Hudson",
	},
	"New Jersey": {
		"Newark",
		"Jersey City",
		"Paterson",
		"Elizabeth",
		"Clifton",
		"Trenton",
		"Camden",
		"Passaic",
		"Union City",
		"Bayonne",
	},
	"New Mexico": {
		"Albuquerque",
		"Las Cruces",
		"Santa Fe",
		"Roswell",
		"Farmington",
		"Rio Rancho",
		"South Valley",
		"Hobbs",
		"Los Lunas",
		"Clovis",
	},
	"New York": {
		"New York City",
		"Buffalo",
		"Rochester",
		"Yonkers",
		"Syracuse",
		"Albany",
		"New Rochelle",
		"Mount Vernon",
		"Schenectady",
		"Utica",
	},
	"North Carolina": {
		"Charlotte",
		"Raleigh",
		"Greensboro",
		"Durham",
		"Winston-Salem",
		"Fayetteville",
		"Cary",
		"Wilmington",
		"High Point",
		"Asheville",
	},
	"North Dakota": {
		"Fargo",
		"Bismarck",
		"Grand Forks",
		"Minot",
		"West Fargo",
		"Mandan",
		"Dickinson",
		"Jamestown",
		"Williston",
		"Wahpeton",
	},
	"Ohio": {
		"Columbus",
		"Cleveland",
		"Cincinnati",
		"Toledo",
		"Akron",
		"Dayton",
		"Youngstown",
		"Canton",
		"Springfield",
		"Elyria",
	},
	"Oklahoma": {
		"Oklahoma City",
		"Tulsa",
		"Norman",
		"Broken Arrow",
		"Lawton",
		"Edmond",
		"Moore",
		"Midwest City",
		"Enid",
		"Stillwater",
	},
	"Oregon": {
		"Portland",
		"Eugene",
		"Salem",
		"Gresham",
		"Hillsboro",
		"Beaverton",
		"Bend",
		"Medford",
		"Springfield",
		"Corvallis",
	},
	"Pennsylvania": {
		"Philadelphia",
		"Pittsburgh",
		"Allentown",
		"Erie",
		"Reading",
		"Scranton",
		"Bethlehem",
		"Lancaster",
		"Harrisburg",
		"Altoona",
	},
	"Rhode Island": {
		"Providence",
		"Warwick",
		"Cranston",
		"Pawtucket",
		"East Providence",
		"Woonsocket",
		"Cumberland",
		"North Providence",
		"South Kingstown",
		"West Warwick",
	},
	"South Carolina": {
		"Columbia",
		"Charleston",
		"North Charleston",
		"Mount Pleasant",
		"Rock Hill",
		"Greenville",
		"Summerville",
		"Sumter",
		"Goose Creek",
		"Hilton Head Island",
	},
	"South Dakota": {
		"Sioux Falls",
		"Rapid City",
		"Aberdeen",
		"Brookings",
		"Watertown",
		"Mitchell",
		"Yankton",
		"Huron",
		"Pierre",
		"Vermillion",
	},
	"Tennessee": {
		"Memphis",
		"Nashville",
		"Knoxville",
		"Chattanooga",
		"Clarksville",
		"Murfreesboro",
		"Franklin",
		"Jackson",
		"Johnson City",
		"Bartlett",
	},
	"Texas": {
		"Houston",
		"San Antonio",
		"Dallas",
		"Austin",
		"Fort Worth",
		"El Paso",
		"Arlington",
		"Corpus Christi",
		"Plano",
		"Laredo",
	},
	"Utah": {
		"Salt Lake City",
		"West Valley City",
		"Provo",
		"West Jordan",
		"Orem",
		"Sandy",
		"Ogden",
		"St. George",
		"Layton",
		"South Jordan",
	},
	"Vermont": {
		"Burlington",
		"South Burlington",
		"Rutland",
		"Barre",
		"Winooski",
		"Montpelier",
		"St. Albans",
		"Newport",
		"Bennington",
		"Brandon",
	},
	"Virginia": {
		"Virginia Beach",
		"Norfolk",
		"Chesapeake",
		"Richmond",
		"Newport News",
		"Alexandria",
		"Hampton",
		"Roanoke",
		"Portsmouth",
		"Suffolk",
	},
	"Washington": {
		"Seattle",
		"Spokane",
		"Tacoma",
		"Vancouver",
		"Bellevue",
		"Kent",
		"Everett",
		"Renton",
		"Yakima",
		"Federal Way",
	},
	"West Virginia": {
		"Charleston",
		"Huntington",
		"Parkersburg",
		"Morgantown",
		"Wheeling",
		"Weirton",
		"Fairmont",
		"Beckley",
		"Martinsburg",
		"Clarksburg",
	},
	"Wisconsin": {
		"Milwaukee",
		"Madison",
		"Green Bay",
		"Kenosha",
		"Racine",
		"Appleton",
		"Waukesha",
		"Oshkosh",
		"Eau Claire",
		"Janesville",
	},
	"Wyoming": {
		"Cheyenne",
		"Casper",
		"Laramie",
		"Gillette",
		"Rock Springs",
		"Sheridan",
		"Green River",
		"Evanston",
		"Riverton",
		"Cody",
	},
}