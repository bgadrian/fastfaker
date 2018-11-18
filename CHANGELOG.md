# FastFaker change log

### v1.1.2
Optimizations

Sentence generator improved with 95% less memory allocations and 50% faster (CPU), affects:
* Sentence, SentenceAvg
* Paragraph, ParagraphAvg
* HipsterSentence, HipsterSentenceAvg
* HipsterParagraph, HipsterParagraphAvg

More tests and /data functionality.

#### v1.1.1
Optimizations and big bug fix.

Fixes digit 9 missing from randDigit, callers affected: 
* Street and StreetNumber
* HexColor 
* Phone and PhoneFormatted and Person
* Username
* Numerify
* CreditCard and Digit

Optimized URL() memory allocations (~20%)
Updated benchmarks
Simplified randDigit
Fixed travis

#### v1.1.0
Moved all the root *go files into a new package: /faker

More tests and examples.
Replace the image generator (the website was too slow) with http://pipsum.com
Adds basic function ()string alternatives to the ones that have parameters (like ImageURL, Sentence and Password).
* Adds AvatarURL() function
* Adds PasswordFull() function
* Adds DateCurrentYear
* Adds DateCurrentYearStr
* Adds DateStr
* Adds SentenceAvg and HipsterSentenceAvg
* Adds ParagraphAvg and HipsterParagraphAvg

Breaking changes (import path)

#### v1.0.0
Moved /data package logic from the faker.
Extracted the pseudo-random main functions to its on package /randomizer.

No breaking changes.

### v1.0.0
Start from the gofakeit 3.14 version:
* changed the name to FastFaker, updated all files

Huge refactor with breaking changes:
* moved all public functions to methods to a new struct (Faker)
* added 2 working modes: Safe or Fast
* new tests, benchmarks and examples
* small updates