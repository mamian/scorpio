Scorpio -To Be The intelligent investor.
===============
> qer.io是为中国证券和期货市场量身定制的开源量化交易平台，推广Golang、Python在国内量化交易领域的应用。 Scorpio是qer.io推出的策略回测和程序化交易框架。

### How to compile 

#### Install [gyp](https://gyp.gsrc.io/)
```
	git clone https://chromium.googlesource.com/external/gyp
	cd gyp
	python setup.py install

```

#### Generate native IDE project files by [gyp](https://gyp.gsrc.io/)
   
   on Linux: Ubuntu 14.04.3 LTS (GNU/Linux 3.13.0-70-generic x86_64)

```
    $ ./gyp/gyp --depth=. QuantTrader.gyp -Dtarget_arch=x64
```
    on windows: Windows 10 with visual studio 2013 express
```
    $ gyp\gyp.bat --depth=. QuantTrader.gyp -Dtarget_arch=x64 -G msvs_version=2013e
```

### set Golang 

Install go linux version

```
	tar -C /usr/local -xzf go1.5.1.linux-amd64.tar.gz
```

Add the folloowing lines to /etc/profile:

```
	#set Go environment
	export GOPATH="/workspace/GoProject:/workspace/QuantProject/qerio.scorpio.v0.2"
	export GOROOT="/usr/local/go"
	export PATH=$GOROOT/bin:$PATH
```

### build SWIG

```
	git clone https://github.com/swig/swig.git
	./autogen.sh 
	apt-get install libpcre++-dev
	apt-get install byacc
	apt-get install yodl
	./configure 
	make 
	make install
```

### build libssl (need by Femas)

```
	$ git clone git://git.openssl.org/openssl.git
	$ ./config
	$ make
	$ make test
	$ make install

	or
	
	$ apt-get install libssl0.9.8   // this works 

```

### Build && Run
```
    cd /workspace/QuantProject/qerio.scorpio.v0.2
	./build.sh && cd src/ && ./scorpio -c 
```

