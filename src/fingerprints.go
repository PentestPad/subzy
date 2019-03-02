package src

type Fingerprint struct {
	engine        string
	status        string
	fingerprint   string
	discussion    string
	documentation string
}

func Fingerprints() [22]Fingerprint {

	var fingerprints [22]Fingerprint

	fingerprints[0] = Fingerprint{
		"AWS/S3",
		"Vulnerable",
		"The specified bucket does not exist",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/36",
		"Not available",
	}

	fingerprints[1] = Fingerprint{
		"Bitbucket",
		"Vulnerable",
		"Repository not found",
		"Not available",
		"Not available",
	}

	fingerprints[2] = Fingerprint{
		"Cloudfront",
		"Edge case",
		"Bad Request: ERROR: The request could not be satisfied",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/29",
		"Not available",
	}

	fingerprints[3] = Fingerprint{
		"Desk",
		"Not vulnerable",
		"Please try again or try Desk.com free for 14 days.",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/9",
		"Not available",
	}

	fingerprints[4] = Fingerprint{
		"Fastly",
		"Edge case",
		"Fastly error: unknown domain:",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/22",
		"Not available",
	}

	fingerprints[5] = Fingerprint{
		"Feedpress",
		"Vulnerable",
		"The feed has not been found.",
		"https://hackerone.com/reports/195350",
		"Not available",
	}

	fingerprints[6] = Fingerprint{
		"Ghost",
		"Vulnerable",
		"The thing you were looking for is no longer here, or never was",
		"",
		"Not available",
	}

	fingerprints[7] = Fingerprint{
		"Github",
		"Vulnerable",
		"There isn't a Github Pages site here",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/37",
		"Not available",
	}

	fingerprints[8] = Fingerprint{
		"Help Juice",
		"Vulnerable",
		"We could not find what you're looking for",
		"Not available",
		"https://help.helpjuice.com/34339-getting-started/custom-domain",
	}

	fingerprints[9] = Fingerprint{
		"Help Scout",
		"Vulnerable",
		"No settings were found for this company",
		"Not available",
		"https://docs.helpscout.net/article/42-setup-custom-domain",
	}

	fingerprints[10] = Fingerprint{
		"Heroku",
		"Edge case",
		"No such app",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/38",
		"Not available",
	}

	fingerprints[11] = Fingerprint{
		"JetBrains",
		"Vulnerable",
		"is not a registered InCloud YouTrack",
		"Not available",
		"Not available",
	}

	fingerprints[12] = Fingerprint{
		"Mashery",
		"Not vulnerable",
		"Unrecognized domain",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/14",
		"Not available",
	}

	fingerprints[13] = Fingerprint{
		"Readme.io",
		"Vulnerable",
		"Project doesnt exist... yet!",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/41",
		"Not available",
	}

	fingerprints[14] = Fingerprint{
		"Shopify",
		"Edge Case",
		"Sorry, this shop is currently unavailable",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/32",
		"https://medium.com/@thebuckhacker/how-to-do-55-000-subdomain-takeover-in-a-blink-of-an-eye-a94954c3fc75",
	}

	fingerprints[15] = Fingerprint{
		"Surge.sh",
		"Vulnerable",
		"project not found",
		"Not available",
		"https://surge.sh/help/adding-a-custom-domain",
	}

	fingerprints[16] = Fingerprint{
		"Tumblr",
		"Vulnerable",
		"Whatever you were looking for doesn't currently exist at this address",
		"Not available",
		"Not available",
	}

	fingerprints[17] = Fingerprint{
		"Tilda",
		"Edge Case",
		"Please renew your subscription",
		"https://github.com/EdOverflow/can-i-take-over-xyz/pull/20",
		"Not available",
	}

	fingerprints[18] = Fingerprint{
		"Unbounce",
		"Not vulnerable",
		"The requested URL was not found on this server",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/11",
		"Not available",
	}

	fingerprints[19] = Fingerprint{
		"UserVoice",
		"Vulnerable",
		"This UserVoice subdomain is currently available",
		"Not available",
		"Not available",
	}

	fingerprints[20] = Fingerprint{
		"Wordpress",
		"Vulnerable",
		"Do you want to register ",
		"https://hackerone.com/reports/274336",
		"Not available",
	}

	fingerprints[21] = Fingerprint{
		"Zendesk",
		"Not Vulnerable",
		"Help Center Closed",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/23",
		"https://support.zendesk.com/hc/en-us/articles/203664356-Changing-the-address-of-your-Help-Center-subdomain-host-mapping-",
	}

	return fingerprints
}
