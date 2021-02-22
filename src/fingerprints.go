package src

type Fingerprint struct {
	engine        string
	status        string
	fingerprint   string
	discussion    string
	documentation string
}

func Fingerprints() []Fingerprint {

	var fingerprints []Fingerprint

	fingerprints = append(fingerprints, Fingerprint{
		"AWS/S3",
		"Vulnerable",
		"The specified bucket does not exist",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/36",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Bitbucket",
		"Vulnerable",
		"Repository not found",
		"Not available",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Cloudfront",
		"Edge case",
		"Bad Request: ERROR: The request could not be satisfied",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/29",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Desk",
		"Not vulnerable",
		"Please try again or try Desk.com free for 14 days.",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/9",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Fastly",
		"Edge case",
		"Fastly error: unknown domain:",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/22",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Feedpress",
		"Vulnerable",
		"The feed has not been found.",
		"https://hackerone.com/reports/195350",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Ghost",
		"Vulnerable",
		"The thing you were looking for is no longer here, or never was",
		"",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Github",
		"Vulnerable",
		"There isn't a Github Pages site here",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/37",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Help Juice",
		"Vulnerable",
		"We could not find what you're looking for",
		"Not available",
		"https://help.helpjuice.com/34339-getting-started/custom-domain",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Help Scout",
		"Vulnerable",
		"No settings were found for this company",
		"Not available",
		"https://docs.helpscout.net/article/42-setup-custom-domain",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Heroku",
		"Edge case",
		"No such app",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/38",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"JetBrains",
		"Vulnerable",
		"is not a registered InCloud YouTrack",
		"Not available",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Mashery",
		"Not vulnerable",
		"Unrecognized domain",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/14",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Readme.io",
		"Vulnerable",
		"Project doesnt exist... yet!",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/41",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Shopify",
		"Edge Case",
		"Sorry, this shop is currently unavailable",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/32",
		"https://medium.com/@thebuckhacker/how-to-do-55-000-subdomain-takeover-in-a-blink-of-an-eye-a94954c3fc75",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Surge.sh",
		"Vulnerable",
		"project not found",
		"Not available",
		"https://surge.sh/help/adding-a-custom-domain",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Tumblr",
		"Vulnerable",
		"Whatever you were looking for doesn't currently exist at this address",
		"Not available",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Tilda",
		"Edge Case",
		"Please renew your subscription",
		"https://github.com/EdOverflow/can-i-take-over-xyz/pull/20",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Unbounce",
		"Not vulnerable",
		"The requested URL was not found on this server",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/11",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"UserVoice",
		"Vulnerable",
		"This UserVoice subdomain is currently available",
		"Not available",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Wordpress",
		"Vulnerable",
		"Do you want to register ",
		"https://hackerone.com/reports/274336",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Zendesk",
		"Not Vulnerable",
		"Help Center Closed",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/23",
		"https://support.zendesk.com/hc/en-us/articles/203664356-Changing-the-address-of-your-Help-Center-subdomain-host-mapping-",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Acquia",
		"Not vulnerable",
		"Web Site Not Found",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/103",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Agile CRM",
		"Vulnerable",
		"Sorry, this page is no longer available.",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/145",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Airee.ru",
		"Vulnerable",
		"Ошибка 402. Сервис",
		"Not available",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Anima",
		"Vulnerable",
		"If this is your website and you've just created it, try refreshing in a minute",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/126",
		"https://docs.animaapp.com/v1/launchpad/08-custom-domain.html",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Campaign Monitor",
		"Vulnerable",
		"Trying to access your account?",
		"Not available",
		"https://help.campaignmonitor.com/custom-domain-names",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Digital Ocean",
		"Vulnerable",
		"Domain uses DO name serves with no records in DO.",
		"Not available",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Gemfury",
		"Vulnerable",
		"404: This page could not be found.",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/154",
		"https://khaledibnalwalid.wordpress.com/2020/06/25/gemfury-subdomain-takeover/",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Google Cloud Storage",
		"Not vulnerable",
		"The specified bucket does not exist.",
		"Not available",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"HatenaBlog",
		"Vulnerable",
		"404 Blog is not found",
		"Not available",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Intercom",
		"Vulnerable",
		"Uh oh. That page doesn't exist.",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/69",
		"https://www.intercom.com/help/",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Kinsta",
		"Vulnerable",
		"No Site For Domain",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/48",
		"https://kinsta.com/knowledgebase/add-domain/",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"LaunchRock",
		"Vulnerable",
		"It looks like you may have taken a wrong turn somewhere. Don't worry...it happens to all of us.",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/74",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Ngrok",
		"Vulnerable",
		"ngrok.io not found",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/92",
		"https://ngrok.com/docs#http-custom-domains",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Pantheon",
		"Vulnerable",
		"404 error unknown site!",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/24",
		"https://medium.com/@hussain_0x3c/hostile-subdomain-takeover-using-pantheon-ebf4ab813111",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Pingdom",
		"Vulnerable",
		"Sorry, couldn't find the status page",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/144",
		"https://help.pingdom.com/hc/en-us/articles/205386171-Public-Status-Page",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"SmartJobBoard",
		"Vulnerable",
		"This job board website is either expired or its domain name is invalid.",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/139",
		"https://help.smartjobboard.com/en/articles/1269655-connecting-a-custom-domain-name",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Smartling",
		"Edge Case",
		"Domain is not configured",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/67",
		"Not available",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Statuspage",
		"Not Vulnerable",
		"Status page pushed a DNS verification in order to prevent malicious takeovers what they mentioned in",
		"https://github.com/EdOverflow/can-i-take-over-xyz/pull/105",
		"https://help.statuspage.io/knowledge_base/topics/domain-ownership",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Uberflip",
		"Vulnerable",
		"Non-hub domain, The URL you've accessed does not provide a hub.",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/150",
		"https://help.uberflip.com/hc/en-us/articles/360018786372-Custom-Domain-Set-up-Your-Hub-on-a-Subdomain",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Webflow",
		"Edge Case",
		"The page you are looking for doesn't exist or has been moved.",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/44",
		"https://forum.webflow.com/t/hosting-a-subdomain-on-webflow/59201",
	})

	fingerprints = append(fingerprints, Fingerprint{
		"Worksites",
		"Vulnerable",
		"Hello! Sorry, but the website you&rsquo;re looking for doesn&rsquo;t exist.",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/142",
		"Not available",
	})

	return fingerprints
}
