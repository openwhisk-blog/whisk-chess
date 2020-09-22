# Whisk Chess

This is `whisk-chess`, the game of chess implemented as a serverless cloud application, using OpenWhisk. This single project may be deployed, as is, to multiple serverless clouds. It is an example of the power of a true, portable, multicloud serverless platform.

Here is a list of deployments in the various cloud:

- [Adobe IO](https://whisk-chess.adobeioruntime.net/api/v1/web/default/chess)
- [IBM Cloud](https://eu-de.functions.appdomain.cloud/api/v1/web/a1d40f6b-e5e3-4f07-8f92-77b525392253/default/chess)
- [Naver](https://wka9bi13u3.apigw.ntruss.com/chess/chess/ZC2o7bFh0x/http)
- [Nimbella](https://apigcp.nimbella.io/api/v1/web/msciabar-zc3thebgxgh/default/chess)

You can read more on [this blog post](https://openwhisk.blog/portability).

You can deploy your own serverless chess app to the [Nimbella serverless cloud](https://nimbella.com) in three steps. Skip 1 & 2 if you already have an account setup locally on your machine.

1. Download the [Nimbella CLI](https://apigcp.nimbella.io/downloads/nim/nim.html#install-the-nimbella-command-line-tool-nim).
2. Run `nim login` to login or sign-up if you don't have an account (it's free and convenient).
3. **Run the Nimbella `deploy` command.**
```bash
nim project deploy github:openwhisk-blog/whisk-chess
```

You will see output that looks like this:
```
Deploying project 'github:openwhisk-blog/whisk-chess'
  to namespace 'your-example-domain'
  on host 'https://apigcp.nimbella.io'
Started  running ./build.sh in github:openwhisk-blog/whisk-chess/packages/default/chess
Finished running ./build.sh in github:openwhisk-blog/whisk-chess/packages/default/chess

Deployed 1 web content items to
  https://your-example-domain-apigcp.nimbella.io
Deployed actions:
  - chess
```

You can open your browser to the shown web URL to start playing chess. Enjoy!
