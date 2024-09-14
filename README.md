# Forward Proxy

What is forward proxy?
>
> From chatGPT:
>
> A forward proxy is a type of server that acts as an intermediary between a client and the internet. When a client makes a request to access a website or online resource, the request is first sent to the forward proxy server. The proxy then forwards this request to the destination server, retrieves the response, and sends it back to the client.


## Local Setup Guide

### Clone the repository
```
git clone git@github.com:chinmayagrawal775/forward_proxy.git
```

### Go to project folder
```
cd forward_proxy
```

### Run Proxy server using below command
```
go run main.go
```

Above command will run the proxy server at below URL:
> http://127.0.0.1:6969

By default it also starts the profiling server. You can access it at:
> http://127.0.0.1:6060/debug/pprof

### Test Proxy server
Run the below commands in new terminal
```
curl --proxy 127.0.0.1:6969 example.com

curl --proxy 127.0.0.1:6969 facebook.com
```
With the first commnad you will see the html output in your terminal.

And with the second command you will see `Access to the site blocked!!` as `facebook.com` is a blocked site in proxy server. You can tweak it by modifying the `config/restricted-hosts.txt` file.

## How to use it as actual proxy server
Here i am giving the example of using your proxy server in [FireFox Browser](https://www.mozilla.org/en-US/firefox/). You can easily configure your firefox browser to use this forward proxy. Below are the steps to do that:

- Open Firefox, go to settings.
- In the General tab, at the very bottom, you will find the `Network Setting` section. Click the `Settings...` button
- Or you can also serach for `proxy` in the search bar of firefox settings
- In the opened modal, click the `Manual proxy configuration`
- Now you will have access to input the proxy URL.
- In HTTP proxy section enter the URL: `127.0.0.1` PORT: `6969`
- do check the `Also use this proxy for HTTPS` checkbox, so that HTTPs URLs will also be proxied.
- Click on `OK`

Hurray 🎉🎉, Now you have cofigured the forward proxy in your firefox browser. All your request will now go through this proxy server. To verify check out the proxy server logs in terminal. You will find various requests logging there.