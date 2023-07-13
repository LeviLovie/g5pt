# g5pt
It's app for encrypt and decrypt data from file to hexidemal data format + encrypt it.

## Important Information, for all versions
  * The app has own data format (look to relevant chapter);
  * The data format isn't similar on different versions of the app;
  * App works only with ASCII.

## Data Format
The app has it's own data format. The app using it to encrypt your data, in security, 
because if it will use format from different app, and this app can decode it.
And, important, that data format isn't similar on different versions of app.

### Decoding of data format
(for v1.2.0)
Here we will see all path of your data from ```<input file>```.
Let's encrypt file, that's look like this:
```Hello, World```

* App is adding aditional information to line (line number, line lenght, and etc. in future);
  ```Hello, World```  -> ```:000001:000012:Hello, World:```
* App translatimg it into hexidemal form;
  ```:000001:000012:Hello, World:```  -> ```3a3030303030313a3030303031323a48656c6c6f2c20576f726c643a```
* App encrypt it, by it's method.
  ```3a3030303030313a3030303031323a48656c6c6f2c20576f726c643a```  -> ```d5fe99ff28a7eed24024bd2630436709d6a9d02799f1c4c68c3b82e5903e6a18```

Important, that app also adding special mark into first line of file (all line from ```<input file>``` writing after it).
And it's different in different versions (look to relevant chapter in every version)

## v1.0.0
### Usage: 
* ```Os X``` and ```Linux``` (file must be in correct dyrectory): ```./g5pt <mode> <input file> <output file>```. About ```<mode>```:
  * ```e``` or ```encrypt``` for encrypt data from ```<input file>``` to ```<output file>```;
  * ```d``` or ```decrypt``` for decrypt data from ```<input file>``` to ```<output file>```.
* ```Windows``` (file must be in correct dyrectory): ```g5pt <mode> <input file> <output file>```. About ```<mode>```:
  * ```e```, ```encrypt``` for encrypt data from ```<input file>``` to ```<output file>```;
  * ```d```, ```decrypt``` for decrypt data from ```<input file>``` to ```<output file>```.

### Important information
This version can read file of future versions, cause here you are setting, encrypt ir decrypt file, and in second choose, it's just reading from second line.

### Mark in start of file
```0x67 0x35 0x70 0x74 0x20 0x64 0x61 0x74 0x61 0x20 0x66 0x6f 0x72 0x6d 0x61 0x74``` = ```g5pt data format```

## v1.1.0
### Usage:
* ```Os X``` and ```Linux``` (file must be in correct dyrectory): ```./g5pt <input file> <output file>```.
* ```Windows``` (file must be in correct dyrectory): ```g5pt <input file> <output file>```.

### Important information
In that version, I did auto getting mode status, but in this cause, it can't read message from v1.0.0. Because it's looking to first line, and checking, is it have format mark.
I will make it better in v1.2.0 (now it's in developing)

### Mark in start of file
```0x67 0x35 0x70 0x74 0x20 0x64 0x61 0x74 0x61 0x20 0x66 0x6f 0x72 0x6d 0x61 0x74 0x2c 0x20 0x76 0x31 0x2e 0x31 0x2e 0x30``` = ```g5pt data format, v1.1.0```
