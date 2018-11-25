# FastFaker change log

### v2.0.2
Improves TemplateCustom/Template performance, reduces the memory allocations by ~25% by reusing internal structures (object pooling).

### v2.0.1
Fixes templates Regex cache parallel writes panic
 
Adds a more complex example: Analytics 

### v2.0.0
Brand new template system. Unlike the old Generate() and Struct() methods that were limited to the /data and basic types, templates allows custom patterns to call all the Faker methods. 

The method Generate() was removed and replaced with Template(), acts in the same way, but with different variables. For more details see the [Templates readme](./TEMPLATES.md)

Other changes:
* modified all benchmarks to use the FastFaker (instead of Safe)
* new method: Binary ("0" or "1")
* new method: BoolText ("false" or "true")
* new examples, benchmarks and tests
* removed Generate method
* Struct and HackerPhrase now use Template* methods

### v1.1.2
Optimizations

Sentence generator improved with 95% less memory allocations and 50% faster (CPU), affects:
* Sentence, SentenceAvg
* Paragraph, ParagraphAvg
* HipsterSentence, HipsterSentenceAvg
* HipsterParagraph, HipsterParagraphAvg

More tests and /data functionality.

### v1.1.1
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

### v1.1.0
**Breaking changes (import path)**: moved all the root *go files into a new package: /faker

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

### v1.0.0
Moved /data package logic from the faker.

Extracted the pseudo-random main functions to its on package /randomizer.

No breaking changes.

### v1.0.0
Start from the gofakeit 3.14 version:
* changed the name to FastFaker, updated all files

Huge refactor with **breaking changes**:
* moved all public functions to methods to a new struct (Faker)
* added 2 working modes: Safe or Fast
* new tests, benchmarks and examples
* small updates