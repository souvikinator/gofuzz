[![fuzz-removebg-preview.png](https://i.postimg.cc/VsxCTCxS/fuzz-removebg-preview.png)](https://postimg.cc/rz9sRKFc)

<hr />

## What is it?

GOFUZZ is fast web fuzzer which takes in URL as input and test the URL for diffrent set of inputs provided by the user.
Currently in Beta phase (now that sounds professional xD)

![gofuzz in action](https://i.imgur.com/orlvQJX.gif)

**results**:

![gofuzz result](https://i.imgur.com/BDuFc09.png)

ah! so we have some forbidden directories ;)

Output is exported to a file and not displayed on the screen to avoid bloating and filling screen with output.

## TODO

- [x] Add Output file feature where output can be stored in specified file
- [X] Add export type TXT 
- [x] Add export type JSON
- [x] Add exclude option which lets user exclude specific response status codes from the results
- [x] Add percentage/progress feature
- [x] Add timeout feature when one URL is not responding for a specific time
- [x] Add GET method feature  
- [ ] Add redirection URL to the results
- [ ] Make a rate limiter
- [ ] Add export type CSV
- [ ] Add Permuation feature
- [ ] Add POST method feature.

and a lot more... 

Will add as we go along

## Features

### -u (URL)

Target URL has to be provided using `-u` option like so:

```bash
gofuzz -u "http://targeturl.com/targetpath?q1=<@>&q2=<@>"
```
**What is `<@>` ?**

`<@>` is placeholder where the test cases will be placed while fuzzing. We'll see how it works on the way. You can place multiple placeholders in the target URL

### 

### -n (numeric)

Numeric values can be passed using `-n` option like so:

- `-n 100` : tests from 0 to 100
- `-n 10,200` : tests from 10 to 200
- `-n 10,11,20,50` : tests for 10,11,20,50 only

```bash
gofuzz -u "httpL//targeturl.com/targetpath?q1=<@>&q2=<@>" -n 100
```
above tests URL for `2000-3000` replacing placeholders(`<@>`) with numbers. Here is an gif showing example:

<p align="center">
  <img src="https://i.imgur.com/VFO6Z34.gif" />
</p>

and here we have the results

<p align="center">
  <img src="https://i.imgur.com/LWT064D.png" />
</p>


### -a (ASCII)

Suppose I want to test a URL for vulnerabilites like SQL injection or LDAP injection. Common way to do it is test for `*,",',=...so on`. Doing it manually is no cool. Provide a range of ASCII values using `-a` option and rest is done by GOFUZZ.

- `-a 65` : tests for `A` only
- `-a 65,90` : tests from `A` to `Z`
- `-a 65,66,67,68` : tests for `A,B,C,D` only

<p align="center">
   <img src="https://i.imgur.com/FY3eRPh.gif" />
</p>

**Results:**

<p align="center">
   <img src="https://i.imgur.com/2BJnxDW.png" />
</p>

### -c (characters)

You can pass list of characters you want to test for, like so

- `-a "{,},^,%,&,*,#,@,!"` : tests for `{,},^,%,&,*,#,@,!` only

NOTE: it is preffered to wrap the input around quotes as shows above to prevent any ambiguity with the shell symbols.


### -o (output directory)

Takes in output directory where the results will be saved. Default is `./output`.

usage: `gofuzz -u "http://targeturl/targetpath?tq1=<@>&tq2=<@>" -f keywords.txt -o ./custom_output_dir`

### -export (result export type)(default:json)

Takes in **txt** or **json** as input.

usage: `gofuzz -u "http://targeturl/targetpath?tq1=<@>&tq2=<@>" -f keywords.txt -export txt`

### -exclude (blacklisting status code)

Takes in status codes as input and doesn't includes them in the result. Example can be seen in the very first gif of this readme.

### -t  (timeout)(default:30000)

Takes in time in milliseconds(ms). How long gofuzz will wait if the connection is not responding. Default 30000 ms or 30 s

Let's set timeout to 1 min or 60 sec or 60000 ms
usage: `gofuzz -u "http://targeturl/targetpath?tq1=<@>&tq2=<@>" -f keywords.txt -t 60000`

### -h (shows usage menu)

#### more features to be added...
