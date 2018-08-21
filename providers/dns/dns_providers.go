package dns

import (
	"fmt"

	"github.com/icasei/lego/acme"
	"github.com/icasei/lego/providers/dns/acmedns"
	"github.com/icasei/lego/providers/dns/auroradns"
	"github.com/icasei/lego/providers/dns/azure"
	"github.com/icasei/lego/providers/dns/bluecat"
	"github.com/icasei/lego/providers/dns/cloudflare"
	"github.com/icasei/lego/providers/dns/cloudxns"
	"github.com/icasei/lego/providers/dns/digitalocean"
	"github.com/icasei/lego/providers/dns/dnsimple"
	"github.com/icasei/lego/providers/dns/dnsmadeeasy"
	"github.com/icasei/lego/providers/dns/dnspod"
	"github.com/icasei/lego/providers/dns/duckdns"
	"github.com/icasei/lego/providers/dns/dyn"
	"github.com/icasei/lego/providers/dns/exec"
	"github.com/icasei/lego/providers/dns/exoscale"
	"github.com/icasei/lego/providers/dns/fastdns"
	"github.com/icasei/lego/providers/dns/gandi"
	"github.com/icasei/lego/providers/dns/gandiv5"
	"github.com/icasei/lego/providers/dns/gcloud"
	"github.com/icasei/lego/providers/dns/glesys"
	"github.com/icasei/lego/providers/dns/godaddy"
	"github.com/icasei/lego/providers/dns/lightsail"
	"github.com/icasei/lego/providers/dns/linode"
	"github.com/icasei/lego/providers/dns/namecheap"
	"github.com/icasei/lego/providers/dns/namedotcom"
	"github.com/icasei/lego/providers/dns/nifcloud"
	"github.com/icasei/lego/providers/dns/ns1"
	"github.com/icasei/lego/providers/dns/otc"
	"github.com/icasei/lego/providers/dns/ovh"
	"github.com/icasei/lego/providers/dns/pdns"
	"github.com/icasei/lego/providers/dns/rackspace"
	"github.com/icasei/lego/providers/dns/rfc2136"
	"github.com/icasei/lego/providers/dns/route53"
	"github.com/icasei/lego/providers/dns/sakuracloud"
	"github.com/icasei/lego/providers/dns/vegadns"
	"github.com/icasei/lego/providers/dns/vultr"
)

// NewDNSChallengeProviderByName Factory for DNS providers
func NewDNSChallengeProviderByName(name string) (acme.ChallengeProvider, error) {
	switch name {
	case "acme-dns":
		return acmedns.NewDNSProvider()
	case "azure":
		return azure.NewDNSProvider()
	case "auroradns":
		return auroradns.NewDNSProvider()
	case "bluecat":
		return bluecat.NewDNSProvider()
	case "cloudflare":
		return cloudflare.NewDNSProvider()
	case "cloudxns":
		return cloudxns.NewDNSProvider()
	case "digitalocean":
		return digitalocean.NewDNSProvider()
	case "dnsimple":
		return dnsimple.NewDNSProvider()
	case "dnsmadeeasy":
		return dnsmadeeasy.NewDNSProvider()
	case "dnspod":
		return dnspod.NewDNSProvider()
	case "duckdns":
		return duckdns.NewDNSProvider()
	case "dyn":
		return dyn.NewDNSProvider()
	case "fastdns":
		return fastdns.NewDNSProvider()
	case "exoscale":
		return exoscale.NewDNSProvider()
	case "gandi":
		return gandi.NewDNSProvider()
	case "gandiv5":
		return gandiv5.NewDNSProvider()
	case "glesys":
		return glesys.NewDNSProvider()
	case "gcloud":
		return gcloud.NewDNSProvider()
	case "godaddy":
		return godaddy.NewDNSProvider()
	case "lightsail":
		return lightsail.NewDNSProvider()
	case "linode":
		return linode.NewDNSProvider()
	case "manual":
		return acme.NewDNSProviderManual()
	case "namecheap":
		return namecheap.NewDNSProvider()
	case "namedotcom":
		return namedotcom.NewDNSProvider()
	case "nifcloud":
		return nifcloud.NewDNSProvider()
	case "rackspace":
		return rackspace.NewDNSProvider()
	case "route53":
		return route53.NewDNSProvider()
	case "rfc2136":
		return rfc2136.NewDNSProvider()
	case "sakuracloud":
		return sakuracloud.NewDNSProvider()
	case "vultr":
		return vultr.NewDNSProvider()
	case "ovh":
		return ovh.NewDNSProvider()
	case "pdns":
		return pdns.NewDNSProvider()
	case "ns1":
		return ns1.NewDNSProvider()
	case "otc":
		return otc.NewDNSProvider()
	case "exec":
		return exec.NewDNSProvider()
	case "vegadns":
		return vegadns.NewDNSProvider()
	default:
		return nil, fmt.Errorf("unrecognised DNS provider: %s", name)
	}
}
