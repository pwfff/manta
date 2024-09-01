package manta

import (
	"github.com/golang/protobuf/proto"
)

type fieldPatch struct {
	minBuild uint32
	maxBuild uint32
	patch    func(f *field)
}

var fieldPatches = []fieldPatch{
	fieldPatch{0, 990, func(f *field) {
		//fmt.Println("patch 1")
		switch f.varName {
		case
			"angExtraLocalAngles",
			"angLocalAngles",
			"m_angInitialAngles",
			"m_angRotation",
			"m_ragAngles",
			"m_vLightDirection":
			if f.parentName == "CBodyComponentBaseAnimatingOverlay" {
				f.encoder = "qangle_pitch_yaw"
			} else {
				f.encoder = "QAngle"
			}

		case
			"dirPrimary",
			"localSound",
			"m_flElasticity",
			"m_location",
			"m_poolOrigin",
			"m_ragPos",
			"m_vecEndPos",
			"m_vecLadderDir",
			"m_vecPlayerMountPositionBottom",
			"m_vecPlayerMountPositionTop",
			"m_viewtarget",
			"m_WorldMaxs",
			"m_WorldMins",
			"origin",
			"vecLocalOrigin":
			f.encoder = "coord"

		case "m_vecLadderNormal":
			f.encoder = "normal"
		}
	}},
	fieldPatch{0, 954, func(f *field) {
		//fmt.Println("patch 2")
		switch f.varName {
		case "m_flMana", "m_flMaxMana":
			f.lowValue = nil
			f.highValue = proto.Float32(8192.0)
		}
	}},
	fieldPatch{1016, 1027, func(f *field) {
		//fmt.Println("patch 3")
		switch f.varName {
		case
			"m_bItemWhiteList",
			"m_bWorldTreeState",
			"m_iPlayerIDsInControl",
			"m_iPlayerSteamID",
			"m_ulTeamBannerLogo",
			"m_ulTeamBaseLogo",
			"m_ulTeamLogo":
			f.encoder = "fixed64"
		}
	}},
	fieldPatch{0, 0, func(f *field) {
		//fmt.Println("patch 4")
		switch f.varName {
		case "m_flSimulationTime", "m_flAnimTime":
			f.encoder = "simtime"
			//case "m_flRuneTime":
			//	f.encoder = "runetime"
			//	//case "m_vecX", "m_vecY", "m_vecZ":
			//	//	f.encoder = "coord"
			//case "m_flPreSimTime", "m_flStartTime": //, "m_flPreSimTime", "m_flFreezeTransitionDuration":
			//	f.encoder = "runetime"
			//case
			//	"angExtraLocalAngles",
			//	"angLocalAngles",
			//	"m_angInitialAngles",
			//	"m_angRotation",
			//	"m_ragAngles",
			//	"m_vLightDirection":
			//	f.encoder = "qangle_pitch_yaw"
		}
	}},
}

func (p *fieldPatch) shouldApply(build uint32) bool {
	if p.minBuild == 0 && p.maxBuild == 0 {
		return true
	}

	return build >= p.minBuild && build <= p.maxBuild
}
