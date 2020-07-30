---
date: "2020-05-29T00:00:00Z"
published: true
subtitle: Or any other S3-compatible object store
tags:
- programming
- software development
title: Use Terraform S3 Backend With Linode
---

This is mostly a note to my future self since I couldn't find any documentation on using Terraform with
[Linode Object Storage](linode.com) as a Terraform state backend. With that said, I hope to save other intrepid Linode users some
time if they're lucky enough to come across this post. While I've explained the problem in terms of Linode, the solution here applies to any non-S3 object store compatible with the S3 API. Swap in the appropriate `endpoint` for your provider.

Chances are, if you've ever tried using Terraform's S3 backend with a non-S3 object store, you've seen an error similar to the following:

> Error: error using credentials to get account ID: error calling sts:GetCallerIdentity: InvalidClientTokenId: The security token included in the request is invalid.
>        status code: 403, request id: 60a33ac0-e460-4745-89dc-668d72413df3

Linode's Object Storage is S3-compatible, which means anywhere you'd typically use Amazon's S3 object storage -- as long
as you can swap in a custom endpoint URL -- you can use Linode Object Storage instead. Which, by the way, is really cheap.
So you should do it. Think of all the non-Bezos pockets you'll be lining.

## Change Your Backend

It's time to update your Terraform backend. First of all, you can _mostly_ refer to [Terraform's S3
documentation](https://www.terraform.io/docs/backends/types/s3.html) for this, but there are two key changes you'll want
to make.

1. `endpoint`: This is the hostname that Terraform uses to connect to Linode's Object Store API. Terraform thinks
   it's talking to S3 but -- S3-compatibility to the rescue -- we're dropping Linode into place.
2. `skip_credentials_validation`: According to Terraform's documentation: _(Optional) Skip the credentials validation
   via the STS API._ Presumably this is referring to [Amazon
   STS](https://docs.aws.amazon.com/STS), but I'll be honest; I don't care. If this isn't set to true, Terraform will ask
  Amazon if you Linode credentials are valid S3 credentials. You don't want that.

## Example Backend Config

```
terraform {
  backend "s3" {
    bucket = "<YOUR LINODE BUCKET NAME>"
    key    = "/path/to/your/tfstate"
    region = "<LINODE REGION>"                        # e.g. us-east-1
    endpoint = "<LINODE REGION>.linodeobjects.com"    # e.g. us-est-1.linodeobjects.com
    skip_credentials_validation = true                # just do it
    access_key = "<LINODE OBJECT STORE ACCESS KEY>"
    secret_key = "<LINODE OBJECT STORE SECRET ACCESS KEY>"
  }
}
```

That should get you all set up using Linode Object Storage as your Terraform state backend.

Cheers, AC
