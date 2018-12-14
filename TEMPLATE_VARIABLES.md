# FastFaker Templates

This document contains all the available variables, with examples that can be used with the Faker.Template, Faker.TemplateCustom and Faker.Struct methods.

A variable is an alias to the public Faker method with the same name, for example:
* **&#123;name&#125;** will be replaced by the result from Faker.Name() method
* **&#123;hackernoun&#125;** will be replaced by the result from Faker.HackerNoun() method

> Note: the Faker.Template() function has 2 more special characters: '#' and '?', they will be replaced by a digit or an ASCII lowercase letter. If you do want this functionality use Faker.TemplateCustom().

For more info see [TEMPLATES](./TEMPLATES.md).

| Variable | Example |
|---------|--------|
| **&#123;avatarurl&#125;** | http://pipsum.com/80x80.jpg |
| **&#123;beeralcohol&#125;** | 4.7% |
| **&#123;beerblg&#125;** | 9.3°Blg |
| **&#123;beerhop&#125;** | Glacier |
| **&#123;beeribu&#125;** | 42 IBU |
| **&#123;beermalt&#125;** | Victory |
| **&#123;beername&#125;** | Trappistes Rochefort 10 |
| **&#123;beerstyle&#125;** | Strong Ale |
| **&#123;beeryeast&#125;** | 2035 - American Lager |
| **&#123;binary&#125;** | 1 |
| **&#123;booltext&#125;** | false |
| **&#123;browser&#125;** | firefox |
| **&#123;bs&#125;** | plug-and-play |
| **&#123;buzzword&#125;** | Intuitive |
| **&#123;carmaker&#125;** | Skoda |
| **&#123;carmodel&#125;** | Mark Lt |
| **&#123;chromeuseragent&#125;** | Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_9_7) AppleWebKit/5311 (KHTML, like Gecko) Chrome/38.0.807.0 Mobile Safari/5311 |
| **&#123;city&#125;** | Alexaburgh |
| **&#123;color&#125;** | SteelBlue |
| **&#123;company&#125;** | Bauch-Koepp |
| **&#123;companysuffix&#125;** | Group |
| **&#123;country&#125;** | Portugal |
| **&#123;countryabr&#125;** | NC |
| **&#123;creditcardcvv&#125;** | 638 |
| **&#123;creditcardexp&#125;** | 11/27 |
| **&#123;creditcardnumber&#125;** | 6573105628485882 |
| **&#123;creditcardnumberluhn&#125;** | 4120163995833970 |
| **&#123;creditcardtype&#125;** | Discover |
| **&#123;currencylong&#125;** | Georgia Lari |
| **&#123;currencyshort&#125;** | MUR |
| **&#123;datecurrentyearstr&#125;** | 2018-08-30 19:15:01.857757656 +0000 UTC |
| **&#123;datestr&#125;** | 1907-12-09 22:56:18.246794801 +0000 UTC |
| **&#123;day&#125;** | 5 |
| **&#123;digit&#125;** | 9 |
| **&#123;domainname&#125;** | dynamicvortals.net |
| **&#123;domainsuffix&#125;** | name |
| **&#123;email&#125;** | elodygraham@cummings.com |
| **&#123;extension&#125;** | dmp |
| **&#123;firefoxuseragent&#125;** | Mozilla/5.0 (X11; Linux x86_64; rv:7.0) Gecko/1934-22-09 Firefox/37.0 |
| **&#123;firstname&#125;** | Maureen |
| **&#123;float32&#125;** | 3.8818993e+37 |
| **&#123;float64&#125;** | 1.0212889603727035e+308 |
| **&#123;float64unary&#125;** | 0.05966110191321585 |
| **&#123;fueltype&#125;** | Methanol |
| **&#123;gender&#125;** | male |
| **&#123;hackerabbreviation&#125;** | ADP |
| **&#123;hackeradjective&#125;** | 1080p |
| **&#123;hackeringverb&#125;** | copying |
| **&#123;hackernoun&#125;** | interface |
| **&#123;hackerphrase&#125;** | If we override the alarm, we can get to the RSS panel through the neural RSS pixel! |
| **&#123;hackerverb&#125;** | compress |
| **&#123;hexcolor&#125;** | #252o5w |
| **&#123;hipsterparagraphavg&#125;** | Disrupt plaid bitters bicycle rights brooklyn keffiyeh meh bicycle rights flannel wayfarers brooklyn franzen hoodie skateboard mumblecore viral flannel etsy. Pork belly cardigan pinterest asymmetrical 8-bit DIY helvetica readymade pabst cold-pressed flexitarian tousled knausgaard fashion axe cray gastropub you probably haven't heard of them lomo. Heirloom vegan try-hard austin kickstarter gluten-free you probably haven't heard of them lo-fi viral cred tote bag echo mustache leggings kickstarter sustainable viral tilde. Semiotics cold-pressed selfies health celiac polaroid franzen chicharrones keffiyeh 90's occupy carry chicharrones lomo twee chicharrones fixie blog. Neutra iPhone schlitz tacos brunch semiotics before they sold out whatever listicle paleo kickstarter williamsburg poutine meh cornhole scenester plaid seitan. Fixie migas dreamcatcher kombucha tofu biodiesel blue bottle pork belly whatever VHS before they sold out fashion axe shabby chic vegan leggings distillery gentrify loko. Chillwave swag keffiyeh salvia photo booth hella cliche jean shorts portland paleo forage food truck narwhal vice salvia five dollar toast cornhole beard. Put a bird on it VHS gluten-free polaroid five dollar toast paleo cornhole hoodie retro sartorial trust fund flannel ennui tumblr pop-up next level banh mi migas. |
| **&#123;hipstersentenceavg&#125;** | Kogi vinyl swag franzen kombucha paleo pickled tofu Wes Anderson XOXO typewriter mixtape tousled put a bird on it farm-to-table photo booth beard cleanse. |
| **&#123;hipsterword&#125;** | brunch |
| **&#123;hour&#125;** | 9 |
| **&#123;httpmethod&#125;** | GET |
| **&#123;int16&#125;** | -17451 |
| **&#123;int32&#125;** | -620577306 |
| **&#123;int64&#125;** | -1002424676547334720 |
| **&#123;int64positive&#125;** | 8992553915671475604 |
| **&#123;int8&#125;** | 8 |
| **&#123;ipv4address&#125;** | 53.244.179.248 |
| **&#123;ipv6address&#125;** | 2001:cafe:97bf:e486:5fb4:a76d:374:dab4 |
| **&#123;jobdescriptor&#125;** | Human |
| **&#123;joblevel&#125;** | Research |
| **&#123;jobtitle&#125;** | Assistant |
| **&#123;lastname&#125;** | Sawayn |
| **&#123;latitude&#125;** | -50.6167776685567 |
| **&#123;letter&#125;** | w |
| **&#123;longitude&#125;** | 157.6689725184914 |
| **&#123;mimetype&#125;** | application/x-troff-me |
| **&#123;minute&#125;** | 16 |
| **&#123;month&#125;** | July |
| **&#123;name&#125;** | Garry Grimes |
| **&#123;nameprefix&#125;** | Ms. |
| **&#123;namesuffix&#125;** | Jr. |
| **&#123;nanosecond&#125;** | 469716624 |
| **&#123;operauseragent&#125;** | Opera/8.32 (Windows 98; en-US) Presto/2.8.289 Version/12.00 |
| **&#123;paragraphavg&#125;** | Eveniet et eos eos voluptate inventore aliquam ut non consequatur exercitationem ut optio quis ratione quia quos ullam. Fugiat omnis nulla ut sed non et assumenda est ut quidem minima ea ex dolores enim maiores error. Rerum eum sint mollitia at ratione et ut et dolorum error magni autem dolorem deserunt molestiae voluptates excepturi. Quod voluptas eligendi ipsam nihil a quo maiores autem exercitationem nihil illum ipsum autem dolorum est modi architecto. Qui repellendus perferendis rem natus placeat corporis est explicabo asperiores consequatur doloribus error in sint sed nisi quis. Cumque omnis a impedit reiciendis tempore unde dolores distinctio et odit qui et cupiditate odio corrupti et maiores. Placeat ratione dolor necessitatibus voluptatem enim pariatur occaecati perferendis doloremque qui voluptatem qui reprehenderit et molestiae aperiam animi. Non et reiciendis non molestias officiis aut ab molestias incidunt ullam qui impedit quidem corrupti culpa suscipit ut. |
| **&#123;passwordfull&#125;** | st2*$t9QJ$mx |
| **&#123;phone&#125;** | 8736568472 |
| **&#123;phoneformatted&#125;** | 546-901-6786 |
| **&#123;safariuseragent&#125;** | Mozilla/5.0 (Macintosh; PPC Mac OS X 10_5_10 rv:6.0; en-US) AppleWebKit/534.51.7 (KHTML, like Gecko) Version/5.2 Safari/534.51.7 |
| **&#123;safecolor&#125;** | black |
| **&#123;second&#125;** | 4 |
| **&#123;sentenceavg&#125;** | Repellat qui rerum accusamus voluptatem autem molestias delectus incidunt assumenda dolor hic omnis dolores iusto non aliquid quo. |
| **&#123;simplestatuscode&#125;** | 301 |
| **&#123;ssn&#125;** | 276845821 |
| **&#123;state&#125;** | Montana |
| **&#123;stateabr&#125;** | HI |
| **&#123;statuscode&#125;** | 406 |
| **&#123;street&#125;** | 2669 Causewaychester |
| **&#123;streetname&#125;** | Street |
| **&#123;streetnumber&#125;** | 68904 |
| **&#123;streetprefix&#125;** | New |
| **&#123;streetsuffix&#125;** | mouth |
| **&#123;timezone&#125;** | Kamchatka Standard Time |
| **&#123;timezoneabv&#125;** | CAST |
| **&#123;timezonefull&#125;** | (UTC+07:00) Bangkok, Hanoi, Jakarta |
| **&#123;timezoneoffset&#125;** | 5.5 |
| **&#123;transmissiongeartype&#125;** | Manual |
| **&#123;uint16&#125;** | 38397 |
| **&#123;uint32&#125;** | 50909631 |
| **&#123;uint64&#125;** | 6153500303096584387 |
| **&#123;uint8&#125;** | 198 |
| **&#123;url&#125;** | http://www.futureplug-and-play.io/supply-chains/b2b/experiences/real-time |
| **&#123;useragent&#125;** | Opera/8.55 (Windows NT 5.2; en-US) Presto/2.13.219 Version/13.00 |
| **&#123;username&#125;** | Keebler2185 |
| **&#123;uuid&#125;** | 54df69cc-20e9-4f87-8969-326d01e80246 |
| **&#123;vehicletype&#125;** | Passenger car medium |
| **&#123;weekday&#125;** | Monday |
| **&#123;word&#125;** | autem |
| **&#123;year&#125;** | 1964 |
| **&#123;zip&#125;** | 81431 |



> This file is generated automatically at build.

