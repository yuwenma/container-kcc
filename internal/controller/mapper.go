package controller

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/mappings"

	krm "github.com/yuwenma/container-kcc/api/v1"
	pb "github.com/yuwenma/container-kcc/generated/google/container/v1/containerpb"
)

var keyMapping = mappings.NewMapping(&pb.Cluster{}, &krm.Cluster{},
	Spec("displayName"),
	Spec("restrictions"),
	Status("uid"),
	Ignore("createTime"),
	Ignore("updateTime"),
	Ignore("deleteTime"),
	Ignore("etag"),
	Ignore("name"),        // TODO: Should be ResourceID?
	Ignore("annotations"), // TODO: Should not ignore
).
	MapNested(&pb.Restrictions{}, &krm.KeyRestrictions{}, "apiTargets",
		"androidKeyRestrictions", "browserKeyRestrictions", "iosKeyRestrictions", "serverKeyRestrictions").
	MapNested(&pb.AndroidKeyRestrictions{}, &krm.KeyAndroidKeyRestrictions{}, "allowedApplications").
	MapNested(&pb.AndroidApplication{}, &krm.KeyAllowedApplications{}, "packageName", "sha1Fingerprint").
	MapNested(&pb.ApiTarget{}, &krm.KeyApiTargets{}, "methods", "service").
	MapNested(&pb.BrowserKeyRestrictions{}, &krm.KeyBrowserKeyRestrictions{}, "allowedReferrers").
	MapNested(&pb.IosKeyRestrictions{}, &krm.KeyIosKeyRestrictions{}, "allowedBundleIds").
	MapNested(&pb.ServerKeyRestrictions{}, &krm.KeyServerKeyRestrictions{}, "allowedIps").
	MustBuild()
