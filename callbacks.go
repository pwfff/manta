package manta

import (
	"github.com/dotabuff/manta/dota"
	"github.com/golang/protobuf/proto"
)

// Callbacks decodes and routes replay events to callback functions
type Callbacks struct {
	onCDemoStop                                        []func(*dota.CDemoStop) error
	onCDemoFileHeader                                  []func(*dota.CDemoFileHeader) error
	onCDemoFileInfo                                    []func(*dota.CDemoFileInfo) error
	onCDemoSyncTick                                    []func(*dota.CDemoSyncTick) error
	onCDemoSendTables                                  []func(*dota.CDemoSendTables) error
	onCDemoClassInfo                                   []func(*dota.CDemoClassInfo) error
	onCDemoStringTables                                []func(*dota.CDemoStringTables) error
	onCDemoPacket                                      []func(*dota.CDemoPacket) error
	onCDemoSignonPacket                                []func(*dota.CDemoPacket) error
	onCDemoConsoleCmd                                  []func(*dota.CDemoConsoleCmd) error
	onCDemoCustomData                                  []func(*dota.CDemoCustomData) error
	onCDemoCustomDataCallbacks                         []func(*dota.CDemoCustomDataCallbacks) error
	onCDemoUserCmd                                     []func(*dota.CDemoUserCmd) error
	onCDemoFullPacket                                  []func(*dota.CDemoFullPacket) error
	onCDemoSaveGame                                    []func(*dota.CDemoSaveGame) error
	onCDemoSpawnGroups                                 []func(*dota.CDemoSpawnGroups) error
	onCDemoAnimationData                               []func(*dota.CDemoAnimationData) error
	onCDemoAnimationHeader                             []func(*dota.CDemoAnimationHeader) error
	onCNETMsg_NOP                                      []func(*dota.CNETMsg_NOP) error
	onCNETMsg_SplitScreenUser                          []func(*dota.CNETMsg_SplitScreenUser) error
	onCNETMsg_Tick                                     []func(*dota.CNETMsg_Tick) error
	onCNETMsg_StringCmd                                []func(*dota.CNETMsg_StringCmd) error
	onCNETMsg_SetConVar                                []func(*dota.CNETMsg_SetConVar) error
	onCNETMsg_SignonState                              []func(*dota.CNETMsg_SignonState) error
	onCNETMsg_SpawnGroup_Load                          []func(*dota.CNETMsg_SpawnGroup_Load) error
	onCNETMsg_SpawnGroup_ManifestUpdate                []func(*dota.CNETMsg_SpawnGroup_ManifestUpdate) error
	onCNETMsg_SpawnGroup_SetCreationTick               []func(*dota.CNETMsg_SpawnGroup_SetCreationTick) error
	onCNETMsg_SpawnGroup_Unload                        []func(*dota.CNETMsg_SpawnGroup_Unload) error
	onCNETMsg_SpawnGroup_LoadCompleted                 []func(*dota.CNETMsg_SpawnGroup_LoadCompleted) error
	onCNETMsg_DebugOverlay                             []func(*dota.CNETMsg_DebugOverlay) error
	onCSVCMsg_ServerInfo                               []func(*dota.CSVCMsg_ServerInfo) error
	onCSVCMsg_FlattenedSerializer                      []func(*dota.CSVCMsg_FlattenedSerializer) error
	onCSVCMsg_ClassInfo                                []func(*dota.CSVCMsg_ClassInfo) error
	onCSVCMsg_SetPause                                 []func(*dota.CSVCMsg_SetPause) error
	onCSVCMsg_CreateStringTable                        []func(*dota.CSVCMsg_CreateStringTable) error
	onCSVCMsg_UpdateStringTable                        []func(*dota.CSVCMsg_UpdateStringTable) error
	onCSVCMsg_VoiceInit                                []func(*dota.CSVCMsg_VoiceInit) error
	onCSVCMsg_VoiceData                                []func(*dota.CSVCMsg_VoiceData) error
	onCSVCMsg_Print                                    []func(*dota.CSVCMsg_Print) error
	onCSVCMsg_Sounds                                   []func(*dota.CSVCMsg_Sounds) error
	onCSVCMsg_SetView                                  []func(*dota.CSVCMsg_SetView) error
	onCSVCMsg_ClearAllStringTables                     []func(*dota.CSVCMsg_ClearAllStringTables) error
	onCSVCMsg_CmdKeyValues                             []func(*dota.CSVCMsg_CmdKeyValues) error
	onCSVCMsg_BSPDecal                                 []func(*dota.CSVCMsg_BSPDecal) error
	onCSVCMsg_SplitScreen                              []func(*dota.CSVCMsg_SplitScreen) error
	onCSVCMsg_PacketEntities                           []func(*dota.CSVCMsg_PacketEntities) error
	onCSVCMsg_Prefetch                                 []func(*dota.CSVCMsg_Prefetch) error
	onCSVCMsg_Menu                                     []func(*dota.CSVCMsg_Menu) error
	onCSVCMsg_GetCvarValue                             []func(*dota.CSVCMsg_GetCvarValue) error
	onCSVCMsg_StopSound                                []func(*dota.CSVCMsg_StopSound) error
	onCSVCMsg_PeerList                                 []func(*dota.CSVCMsg_PeerList) error
	onCSVCMsg_PacketReliable                           []func(*dota.CSVCMsg_PacketReliable) error
	onCSVCMsg_HLTVStatus                               []func(*dota.CSVCMsg_HLTVStatus) error
	onCSVCMsg_ServerSteamID                            []func(*dota.CSVCMsg_ServerSteamID) error
	onCSVCMsg_FullFrameSplit                           []func(*dota.CSVCMsg_FullFrameSplit) error
	onCSVCMsg_RconServerDetails                        []func(*dota.CSVCMsg_RconServerDetails) error
	onCSVCMsg_UserMessage                              []func(*dota.CSVCMsg_UserMessage) error
	onCSVCMsg_Broadcast_Command                        []func(*dota.CSVCMsg_Broadcast_Command) error
	onCSVCMsg_HltvFixupOperatorStatus                  []func(*dota.CSVCMsg_HltvFixupOperatorStatus) error
	onCUserMessageAchievementEvent                     []func(*dota.CUserMessageAchievementEvent) error
	onCUserMessageCloseCaption                         []func(*dota.CUserMessageCloseCaption) error
	onCUserMessageCloseCaptionDirect                   []func(*dota.CUserMessageCloseCaptionDirect) error
	onCUserMessageCurrentTimescale                     []func(*dota.CUserMessageCurrentTimescale) error
	onCUserMessageDesiredTimescale                     []func(*dota.CUserMessageDesiredTimescale) error
	onCUserMessageFade                                 []func(*dota.CUserMessageFade) error
	onCUserMessageGameTitle                            []func(*dota.CUserMessageGameTitle) error
	onCUserMessageHudMsg                               []func(*dota.CUserMessageHudMsg) error
	onCUserMessageHudText                              []func(*dota.CUserMessageHudText) error
	onCUserMessageColoredText                          []func(*dota.CUserMessageColoredText) error
	onCUserMessageRequestState                         []func(*dota.CUserMessageRequestState) error
	onCUserMessageResetHUD                             []func(*dota.CUserMessageResetHUD) error
	onCUserMessageRumble                               []func(*dota.CUserMessageRumble) error
	onCUserMessageSayText                              []func(*dota.CUserMessageSayText) error
	onCUserMessageSayText2                             []func(*dota.CUserMessageSayText2) error
	onCUserMessageSayTextChannel                       []func(*dota.CUserMessageSayTextChannel) error
	onCUserMessageShake                                []func(*dota.CUserMessageShake) error
	onCUserMessageShakeDir                             []func(*dota.CUserMessageShakeDir) error
	onCUserMessageWaterShake                           []func(*dota.CUserMessageWaterShake) error
	onCUserMessageTextMsg                              []func(*dota.CUserMessageTextMsg) error
	onCUserMessageScreenTilt                           []func(*dota.CUserMessageScreenTilt) error
	onCUserMessageVoiceMask                            []func(*dota.CUserMessageVoiceMask) error
	onCUserMessageSendAudio                            []func(*dota.CUserMessageSendAudio) error
	onCUserMessageItemPickup                           []func(*dota.CUserMessageItemPickup) error
	onCUserMessageAmmoDenied                           []func(*dota.CUserMessageAmmoDenied) error
	onCUserMessageShowMenu                             []func(*dota.CUserMessageShowMenu) error
	onCUserMessageCreditsMsg                           []func(*dota.CUserMessageCreditsMsg) error
	onCEntityMessagePlayJingle                         []func(*dota.CEntityMessagePlayJingle) error
	onCEntityMessageScreenOverlay                      []func(*dota.CEntityMessageScreenOverlay) error
	onCEntityMessageRemoveAllDecals                    []func(*dota.CEntityMessageRemoveAllDecals) error
	onCEntityMessagePropagateForce                     []func(*dota.CEntityMessagePropagateForce) error
	onCEntityMessageDoSpark                            []func(*dota.CEntityMessageDoSpark) error
	onCEntityMessageFixAngle                           []func(*dota.CEntityMessageFixAngle) error
	onCUserMessageCloseCaptionPlaceholder              []func(*dota.CUserMessageCloseCaptionPlaceholder) error
	onCUserMessageCameraTransition                     []func(*dota.CUserMessageCameraTransition) error
	onCUserMessageAudioParameter                       []func(*dota.CUserMessageAudioParameter) error
	onCUserMessageHapticsManagerPulse                  []func(*dota.CUserMessageHapticsManagerPulse) error
	onCUserMessageHapticsManagerEffect                 []func(*dota.CUserMessageHapticsManagerEffect) error
	onCUserMessageUpdateCssClasses                     []func(*dota.CUserMessageUpdateCssClasses) error
	onCUserMessageServerFrameTime                      []func(*dota.CUserMessageServerFrameTime) error
	onCUserMessageLagCompensationError                 []func(*dota.CUserMessageLagCompensationError) error
	onCUserMessageRequestDllStatus                     []func(*dota.CUserMessageRequestDllStatus) error
	onCUserMessageRequestUtilAction                    []func(*dota.CUserMessageRequestUtilAction) error
	onCUserMessageRequestInventory                     []func(*dota.CUserMessageRequestInventory) error
	onCUserMessageRequestDiagnostic                    []func(*dota.CUserMessageRequestDiagnostic) error
	onCMsgVDebugGameSessionIDEvent                     []func(*dota.CMsgVDebugGameSessionIDEvent) error
	onCMsgPlaceDecalEvent                              []func(*dota.CMsgPlaceDecalEvent) error
	onCMsgClearWorldDecalsEvent                        []func(*dota.CMsgClearWorldDecalsEvent) error
	onCMsgClearEntityDecalsEvent                       []func(*dota.CMsgClearEntityDecalsEvent) error
	onCMsgClearDecalsForSkeletonInstanceEvent          []func(*dota.CMsgClearDecalsForSkeletonInstanceEvent) error
	onCMsgSource1LegacyGameEventList                   []func(*dota.CMsgSource1LegacyGameEventList) error
	onCMsgSource1LegacyListenEvents                    []func(*dota.CMsgSource1LegacyListenEvents) error
	onCMsgSource1LegacyGameEvent                       []func(*dota.CMsgSource1LegacyGameEvent) error
	onCMsgSosStartSoundEvent                           []func(*dota.CMsgSosStartSoundEvent) error
	onCMsgSosStopSoundEvent                            []func(*dota.CMsgSosStopSoundEvent) error
	onCMsgSosSetSoundEventParams                       []func(*dota.CMsgSosSetSoundEventParams) error
	onCMsgSosSetLibraryStackFields                     []func(*dota.CMsgSosSetLibraryStackFields) error
	onCMsgSosStopSoundEventHash                        []func(*dota.CMsgSosStopSoundEventHash) error
	onCCitadelUserMessage_Damage                       []func(*dota.CCitadelUserMessage_Damage) error
	onCCitadelUserMsg_MapPing                          []func(*dota.CCitadelUserMsg_MapPing) error
	onCCitadelUserMsg_TeamRewards                      []func(*dota.CCitadelUserMsg_TeamRewards) error
	onCCitadelUserMsg_TriggerDamageFlash               []func(*dota.CCitadelUserMsg_TriggerDamageFlash) error
	onCCitadelUserMsg_AbilitiesChanged                 []func(*dota.CCitadelUserMsg_AbilitiesChanged) error
	onCCitadelUserMsg_RecentDamageSummary              []func(*dota.CCitadelUserMsg_RecentDamageSummary) error
	onCCitadelUserMsg_SpectatorTeamChanged             []func(*dota.CCitadelUserMsg_SpectatorTeamChanged) error
	onCCitadelUserMsg_ChatWheel                        []func(*dota.CCitadelUserMsg_ChatWheel) error
	onCCitadelUserMsg_GoldHistory                      []func(*dota.CCitadelUserMsg_GoldHistory) error
	onCCitadelUserMsg_ChatMsg                          []func(*dota.CCitadelUserMsg_ChatMsg) error
	onCCitadelUserMsg_QuickResponse                    []func(*dota.CCitadelUserMsg_QuickResponse) error
	onCCitadelUserMsg_PostMatchDetails                 []func(*dota.CCitadelUserMsg_PostMatchDetails) error
	onCCitadelUserMsg_ChatEvent                        []func(*dota.CCitadelUserMsg_ChatEvent) error
	onCCitadelUserMsg_AbilityInterrupted               []func(*dota.CCitadelUserMsg_AbilityInterrupted) error
	onCCitadelUserMsg_HeroKilled                       []func(*dota.CCitadelUserMsg_HeroKilled) error
	onCCitadelUserMsg_ReturnIdol                       []func(*dota.CCitadelUserMsg_ReturnIdol) error
	onCCitadelUserMsg_SetClientCameraAngles            []func(*dota.CCitadelUserMsg_SetClientCameraAngles) error
	onCCitadelUserMsg_MapLine                          []func(*dota.CCitadelUserMsg_MapLine) error
	onCCitadelUserMessage_BulletHit                    []func(*dota.CCitadelUserMessage_BulletHit) error
	onCCitadelUserMessage_ObjectiveMask                []func(*dota.CCitadelUserMessage_ObjectiveMask) error
	onCCitadelUserMessage_ModifierApplied              []func(*dota.CCitadelUserMessage_ModifierApplied) error
	onCCitadelUserMsg_CameraController                 []func(*dota.CCitadelUserMsg_CameraController) error
	onCCitadelUserMessage_AuraModifierApplied          []func(*dota.CCitadelUserMessage_AuraModifierApplied) error
	onCCitadelUserMsg_ObstructedShotFired              []func(*dota.CCitadelUserMsg_ObstructedShotFired) error
	onCCitadelUserMsg_AbilityLateFailure               []func(*dota.CCitadelUserMsg_AbilityLateFailure) error
	onCCitadelUserMsg_AbilityPing                      []func(*dota.CCitadelUserMsg_AbilityPing) error
	onCCitadelUserMsg_PostProcessingAnim               []func(*dota.CCitadelUserMsg_PostProcessingAnim) error
	onCCitadelUserMsg_DeathReplayData                  []func(*dota.CCitadelUserMsg_DeathReplayData) error
	onCCitadelUserMsg_PlayerLifetimeStatInfo           []func(*dota.CCitadelUserMsg_PlayerLifetimeStatInfo) error
	onCCitadelUserMsg_ForceShopClosed                  []func(*dota.CCitadelUserMsg_ForceShopClosed) error
	onCCitadelUserMsg_StaminaDrained                   []func(*dota.CCitadelUserMsg_StaminaDrained) error
	onCCitadelUserMessage_AbilityNotify                []func(*dota.CCitadelUserMessage_AbilityNotify) error
	onCCitadelUserMsg_GetDamageStatsResponse           []func(*dota.CCitadelUserMsg_GetDamageStatsResponse) error
	onCCitadelUserMsg_ParticipantStartSoundEvent       []func(*dota.CCitadelUserMsg_ParticipantStartSoundEvent) error
	onCCitadelUserMsg_ParticipantStopSoundEvent        []func(*dota.CCitadelUserMsg_ParticipantStopSoundEvent) error
	onCCitadelUserMsg_ParticipantStopSoundEventHash    []func(*dota.CCitadelUserMsg_ParticipantStopSoundEventHash) error
	onCCitadelUserMsg_ParticipantSetSoundEventParams   []func(*dota.CCitadelUserMsg_ParticipantSetSoundEventParams) error
	onCCitadelUserMsg_ParticipantSetLibraryStackFields []func(*dota.CCitadelUserMsg_ParticipantSetLibraryStackFields) error
	onCCitadelUserMsg_BossKilled                       []func(*dota.CCitadelUserMsg_BossKilled) error
	onCMsgFireBullets                                  []func(*dota.CMsgFireBullets) error
	onCMsgPlayerAnimEvent                              []func(*dota.CMsgPlayerAnimEvent) error
	onCMsgParticleSystemManager                        []func(*dota.CMsgParticleSystemManager) error
	onCMsgScreenTextPretty                             []func(*dota.CMsgScreenTextPretty) error
	onCMsgServerRequestedTracer                        []func(*dota.CMsgServerRequestedTracer) error
	onCMsgBulletImpact                                 []func(*dota.CMsgBulletImpact) error
	onCMsgEnableSatVolumesEvent                        []func(*dota.CMsgEnableSatVolumesEvent) error
	onCMsgPlaceSatVolumeEvent                          []func(*dota.CMsgPlaceSatVolumeEvent) error
	onCMsgDisableSatVolumesEvent                       []func(*dota.CMsgDisableSatVolumesEvent) error
	onCMsgRemoveSatVolumeEvent                         []func(*dota.CMsgRemoveSatVolumeEvent) error

	pb *proto.Buffer
}

func newCallbacks() *Callbacks {
	return &Callbacks{
		pb: &proto.Buffer{},
	}
}

// OnCDemoStop registers a callback EDemoCommands_DEM_Stop
func (c *Callbacks) OnCDemoStop(fn func(*dota.CDemoStop) error) {
	c.onCDemoStop = append(c.onCDemoStop, fn)
}

// OnCDemoFileHeader registers a callback EDemoCommands_DEM_FileHeader
func (c *Callbacks) OnCDemoFileHeader(fn func(*dota.CDemoFileHeader) error) {
	c.onCDemoFileHeader = append(c.onCDemoFileHeader, fn)
}

// OnCDemoFileInfo registers a callback EDemoCommands_DEM_FileInfo
func (c *Callbacks) OnCDemoFileInfo(fn func(*dota.CDemoFileInfo) error) {
	c.onCDemoFileInfo = append(c.onCDemoFileInfo, fn)
}

// OnCDemoSyncTick registers a callback EDemoCommands_DEM_SyncTick
func (c *Callbacks) OnCDemoSyncTick(fn func(*dota.CDemoSyncTick) error) {
	c.onCDemoSyncTick = append(c.onCDemoSyncTick, fn)
}

// OnCDemoSendTables registers a callback EDemoCommands_DEM_SendTables
func (c *Callbacks) OnCDemoSendTables(fn func(*dota.CDemoSendTables) error) {
	c.onCDemoSendTables = append(c.onCDemoSendTables, fn)
}

// OnCDemoClassInfo registers a callback EDemoCommands_DEM_ClassInfo
func (c *Callbacks) OnCDemoClassInfo(fn func(*dota.CDemoClassInfo) error) {
	c.onCDemoClassInfo = append(c.onCDemoClassInfo, fn)
}

// OnCDemoStringTables registers a callback EDemoCommands_DEM_StringTables
func (c *Callbacks) OnCDemoStringTables(fn func(*dota.CDemoStringTables) error) {
	c.onCDemoStringTables = append(c.onCDemoStringTables, fn)
}

// OnCDemoPacket registers a callback EDemoCommands_DEM_Packet
func (c *Callbacks) OnCDemoPacket(fn func(*dota.CDemoPacket) error) {
	c.onCDemoPacket = append(c.onCDemoPacket, fn)
}

// OnCDemoSignonPacket registers a callback EDemoCommands_DEM_SignonPacket
func (c *Callbacks) OnCDemoSignonPacket(fn func(*dota.CDemoPacket) error) {
	c.onCDemoSignonPacket = append(c.onCDemoSignonPacket, fn)
}

// OnCDemoConsoleCmd registers a callback EDemoCommands_DEM_ConsoleCmd
func (c *Callbacks) OnCDemoConsoleCmd(fn func(*dota.CDemoConsoleCmd) error) {
	c.onCDemoConsoleCmd = append(c.onCDemoConsoleCmd, fn)
}

// OnCDemoCustomData registers a callback EDemoCommands_DEM_CustomData
func (c *Callbacks) OnCDemoCustomData(fn func(*dota.CDemoCustomData) error) {
	c.onCDemoCustomData = append(c.onCDemoCustomData, fn)
}

// OnCDemoCustomDataCallbacks registers a callback EDemoCommands_DEM_CustomDataCallbacks
func (c *Callbacks) OnCDemoCustomDataCallbacks(fn func(*dota.CDemoCustomDataCallbacks) error) {
	c.onCDemoCustomDataCallbacks = append(c.onCDemoCustomDataCallbacks, fn)
}

// OnCDemoUserCmd registers a callback EDemoCommands_DEM_UserCmd
func (c *Callbacks) OnCDemoUserCmd(fn func(*dota.CDemoUserCmd) error) {
	c.onCDemoUserCmd = append(c.onCDemoUserCmd, fn)
}

// OnCDemoFullPacket registers a callback EDemoCommands_DEM_FullPacket
func (c *Callbacks) OnCDemoFullPacket(fn func(*dota.CDemoFullPacket) error) {
	c.onCDemoFullPacket = append(c.onCDemoFullPacket, fn)
}

// OnCDemoSaveGame registers a callback EDemoCommands_DEM_SaveGame
func (c *Callbacks) OnCDemoSaveGame(fn func(*dota.CDemoSaveGame) error) {
	c.onCDemoSaveGame = append(c.onCDemoSaveGame, fn)
}

// OnCDemoSpawnGroups registers a callback EDemoCommands_DEM_SpawnGroups
func (c *Callbacks) OnCDemoSpawnGroups(fn func(*dota.CDemoSpawnGroups) error) {
	c.onCDemoSpawnGroups = append(c.onCDemoSpawnGroups, fn)
}

// OnCDemoAnimationData registers a callback EDemoCommands_DEM_AnimationData
func (c *Callbacks) OnCDemoAnimationData(fn func(*dota.CDemoAnimationData) error) {
	c.onCDemoAnimationData = append(c.onCDemoAnimationData, fn)
}

// OnCDemoAnimationHeader registers a callback EDemoCommands_DEM_AnimationHeader
func (c *Callbacks) OnCDemoAnimationHeader(fn func(*dota.CDemoAnimationHeader) error) {
	c.onCDemoAnimationHeader = append(c.onCDemoAnimationHeader, fn)
}

// OnCNETMsg_NOP registers a callback for NET_Messages_net_NOP
func (c *Callbacks) OnCNETMsg_NOP(fn func(*dota.CNETMsg_NOP) error) {
	c.onCNETMsg_NOP = append(c.onCNETMsg_NOP, fn)
}

// OnCNETMsg_SplitScreenUser registers a callback for NET_Messages_net_SplitScreenUser
func (c *Callbacks) OnCNETMsg_SplitScreenUser(fn func(*dota.CNETMsg_SplitScreenUser) error) {
	c.onCNETMsg_SplitScreenUser = append(c.onCNETMsg_SplitScreenUser, fn)
}

// OnCNETMsg_Tick registers a callback for NET_Messages_net_Tick
func (c *Callbacks) OnCNETMsg_Tick(fn func(*dota.CNETMsg_Tick) error) {
	c.onCNETMsg_Tick = append(c.onCNETMsg_Tick, fn)
}

// OnCNETMsg_StringCmd registers a callback for NET_Messages_net_StringCmd
func (c *Callbacks) OnCNETMsg_StringCmd(fn func(*dota.CNETMsg_StringCmd) error) {
	c.onCNETMsg_StringCmd = append(c.onCNETMsg_StringCmd, fn)
}

// OnCNETMsg_SetConVar registers a callback for NET_Messages_net_SetConVar
func (c *Callbacks) OnCNETMsg_SetConVar(fn func(*dota.CNETMsg_SetConVar) error) {
	c.onCNETMsg_SetConVar = append(c.onCNETMsg_SetConVar, fn)
}

// OnCNETMsg_SignonState registers a callback for NET_Messages_net_SignonState
func (c *Callbacks) OnCNETMsg_SignonState(fn func(*dota.CNETMsg_SignonState) error) {
	c.onCNETMsg_SignonState = append(c.onCNETMsg_SignonState, fn)
}

// OnCNETMsg_SpawnGroup_Load registers a callback for NET_Messages_net_SpawnGroup_Load
func (c *Callbacks) OnCNETMsg_SpawnGroup_Load(fn func(*dota.CNETMsg_SpawnGroup_Load) error) {
	c.onCNETMsg_SpawnGroup_Load = append(c.onCNETMsg_SpawnGroup_Load, fn)
}

// OnCNETMsg_SpawnGroup_ManifestUpdate registers a callback for NET_Messages_net_SpawnGroup_ManifestUpdate
func (c *Callbacks) OnCNETMsg_SpawnGroup_ManifestUpdate(fn func(*dota.CNETMsg_SpawnGroup_ManifestUpdate) error) {
	c.onCNETMsg_SpawnGroup_ManifestUpdate = append(c.onCNETMsg_SpawnGroup_ManifestUpdate, fn)
}

// OnCNETMsg_SpawnGroup_SetCreationTick registers a callback for NET_Messages_net_SpawnGroup_SetCreationTick
func (c *Callbacks) OnCNETMsg_SpawnGroup_SetCreationTick(fn func(*dota.CNETMsg_SpawnGroup_SetCreationTick) error) {
	c.onCNETMsg_SpawnGroup_SetCreationTick = append(c.onCNETMsg_SpawnGroup_SetCreationTick, fn)
}

// OnCNETMsg_SpawnGroup_Unload registers a callback for NET_Messages_net_SpawnGroup_Unload
func (c *Callbacks) OnCNETMsg_SpawnGroup_Unload(fn func(*dota.CNETMsg_SpawnGroup_Unload) error) {
	c.onCNETMsg_SpawnGroup_Unload = append(c.onCNETMsg_SpawnGroup_Unload, fn)
}

// OnCNETMsg_SpawnGroup_LoadCompleted registers a callback for NET_Messages_net_SpawnGroup_LoadCompleted
func (c *Callbacks) OnCNETMsg_SpawnGroup_LoadCompleted(fn func(*dota.CNETMsg_SpawnGroup_LoadCompleted) error) {
	c.onCNETMsg_SpawnGroup_LoadCompleted = append(c.onCNETMsg_SpawnGroup_LoadCompleted, fn)
}

// OnCNETMsg_DebugOverlay registers a callback for NET_Messages_net_DebugOverlay
func (c *Callbacks) OnCNETMsg_DebugOverlay(fn func(*dota.CNETMsg_DebugOverlay) error) {
	c.onCNETMsg_DebugOverlay = append(c.onCNETMsg_DebugOverlay, fn)
}

// OnCSVCMsg_ServerInfo registers a callback for SVC_Messages_svc_ServerInfo
func (c *Callbacks) OnCSVCMsg_ServerInfo(fn func(*dota.CSVCMsg_ServerInfo) error) {
	c.onCSVCMsg_ServerInfo = append(c.onCSVCMsg_ServerInfo, fn)
}

// OnCSVCMsg_FlattenedSerializer registers a callback for SVC_Messages_svc_FlattenedSerializer
func (c *Callbacks) OnCSVCMsg_FlattenedSerializer(fn func(*dota.CSVCMsg_FlattenedSerializer) error) {
	c.onCSVCMsg_FlattenedSerializer = append(c.onCSVCMsg_FlattenedSerializer, fn)
}

// OnCSVCMsg_ClassInfo registers a callback for SVC_Messages_svc_ClassInfo
func (c *Callbacks) OnCSVCMsg_ClassInfo(fn func(*dota.CSVCMsg_ClassInfo) error) {
	c.onCSVCMsg_ClassInfo = append(c.onCSVCMsg_ClassInfo, fn)
}

// OnCSVCMsg_SetPause registers a callback for SVC_Messages_svc_SetPause
func (c *Callbacks) OnCSVCMsg_SetPause(fn func(*dota.CSVCMsg_SetPause) error) {
	c.onCSVCMsg_SetPause = append(c.onCSVCMsg_SetPause, fn)
}

// OnCSVCMsg_CreateStringTable registers a callback for SVC_Messages_svc_CreateStringTable
func (c *Callbacks) OnCSVCMsg_CreateStringTable(fn func(*dota.CSVCMsg_CreateStringTable) error) {
	c.onCSVCMsg_CreateStringTable = append(c.onCSVCMsg_CreateStringTable, fn)
}

// OnCSVCMsg_UpdateStringTable registers a callback for SVC_Messages_svc_UpdateStringTable
func (c *Callbacks) OnCSVCMsg_UpdateStringTable(fn func(*dota.CSVCMsg_UpdateStringTable) error) {
	c.onCSVCMsg_UpdateStringTable = append(c.onCSVCMsg_UpdateStringTable, fn)
}

// OnCSVCMsg_VoiceInit registers a callback for SVC_Messages_svc_VoiceInit
func (c *Callbacks) OnCSVCMsg_VoiceInit(fn func(*dota.CSVCMsg_VoiceInit) error) {
	c.onCSVCMsg_VoiceInit = append(c.onCSVCMsg_VoiceInit, fn)
}

// OnCSVCMsg_VoiceData registers a callback for SVC_Messages_svc_VoiceData
func (c *Callbacks) OnCSVCMsg_VoiceData(fn func(*dota.CSVCMsg_VoiceData) error) {
	c.onCSVCMsg_VoiceData = append(c.onCSVCMsg_VoiceData, fn)
}

// OnCSVCMsg_Print registers a callback for SVC_Messages_svc_Print
func (c *Callbacks) OnCSVCMsg_Print(fn func(*dota.CSVCMsg_Print) error) {
	c.onCSVCMsg_Print = append(c.onCSVCMsg_Print, fn)
}

// OnCSVCMsg_Sounds registers a callback for SVC_Messages_svc_Sounds
func (c *Callbacks) OnCSVCMsg_Sounds(fn func(*dota.CSVCMsg_Sounds) error) {
	c.onCSVCMsg_Sounds = append(c.onCSVCMsg_Sounds, fn)
}

// OnCSVCMsg_SetView registers a callback for SVC_Messages_svc_SetView
func (c *Callbacks) OnCSVCMsg_SetView(fn func(*dota.CSVCMsg_SetView) error) {
	c.onCSVCMsg_SetView = append(c.onCSVCMsg_SetView, fn)
}

// OnCSVCMsg_ClearAllStringTables registers a callback for SVC_Messages_svc_ClearAllStringTables
func (c *Callbacks) OnCSVCMsg_ClearAllStringTables(fn func(*dota.CSVCMsg_ClearAllStringTables) error) {
	c.onCSVCMsg_ClearAllStringTables = append(c.onCSVCMsg_ClearAllStringTables, fn)
}

// OnCSVCMsg_CmdKeyValues registers a callback for SVC_Messages_svc_CmdKeyValues
func (c *Callbacks) OnCSVCMsg_CmdKeyValues(fn func(*dota.CSVCMsg_CmdKeyValues) error) {
	c.onCSVCMsg_CmdKeyValues = append(c.onCSVCMsg_CmdKeyValues, fn)
}

// OnCSVCMsg_BSPDecal registers a callback for SVC_Messages_svc_BSPDecal
func (c *Callbacks) OnCSVCMsg_BSPDecal(fn func(*dota.CSVCMsg_BSPDecal) error) {
	c.onCSVCMsg_BSPDecal = append(c.onCSVCMsg_BSPDecal, fn)
}

// OnCSVCMsg_SplitScreen registers a callback for SVC_Messages_svc_SplitScreen
func (c *Callbacks) OnCSVCMsg_SplitScreen(fn func(*dota.CSVCMsg_SplitScreen) error) {
	c.onCSVCMsg_SplitScreen = append(c.onCSVCMsg_SplitScreen, fn)
}

// OnCSVCMsg_PacketEntities registers a callback for SVC_Messages_svc_PacketEntities
func (c *Callbacks) OnCSVCMsg_PacketEntities(fn func(*dota.CSVCMsg_PacketEntities) error) {
	c.onCSVCMsg_PacketEntities = append(c.onCSVCMsg_PacketEntities, fn)
}

// OnCSVCMsg_Prefetch registers a callback for SVC_Messages_svc_Prefetch
func (c *Callbacks) OnCSVCMsg_Prefetch(fn func(*dota.CSVCMsg_Prefetch) error) {
	c.onCSVCMsg_Prefetch = append(c.onCSVCMsg_Prefetch, fn)
}

// OnCSVCMsg_Menu registers a callback for SVC_Messages_svc_Menu
func (c *Callbacks) OnCSVCMsg_Menu(fn func(*dota.CSVCMsg_Menu) error) {
	c.onCSVCMsg_Menu = append(c.onCSVCMsg_Menu, fn)
}

// OnCSVCMsg_GetCvarValue registers a callback for SVC_Messages_svc_GetCvarValue
func (c *Callbacks) OnCSVCMsg_GetCvarValue(fn func(*dota.CSVCMsg_GetCvarValue) error) {
	c.onCSVCMsg_GetCvarValue = append(c.onCSVCMsg_GetCvarValue, fn)
}

// OnCSVCMsg_StopSound registers a callback for SVC_Messages_svc_StopSound
func (c *Callbacks) OnCSVCMsg_StopSound(fn func(*dota.CSVCMsg_StopSound) error) {
	c.onCSVCMsg_StopSound = append(c.onCSVCMsg_StopSound, fn)
}

// OnCSVCMsg_PeerList registers a callback for SVC_Messages_svc_PeerList
func (c *Callbacks) OnCSVCMsg_PeerList(fn func(*dota.CSVCMsg_PeerList) error) {
	c.onCSVCMsg_PeerList = append(c.onCSVCMsg_PeerList, fn)
}

// OnCSVCMsg_PacketReliable registers a callback for SVC_Messages_svc_PacketReliable
func (c *Callbacks) OnCSVCMsg_PacketReliable(fn func(*dota.CSVCMsg_PacketReliable) error) {
	c.onCSVCMsg_PacketReliable = append(c.onCSVCMsg_PacketReliable, fn)
}

// OnCSVCMsg_HLTVStatus registers a callback for SVC_Messages_svc_HLTVStatus
func (c *Callbacks) OnCSVCMsg_HLTVStatus(fn func(*dota.CSVCMsg_HLTVStatus) error) {
	c.onCSVCMsg_HLTVStatus = append(c.onCSVCMsg_HLTVStatus, fn)
}

// OnCSVCMsg_ServerSteamID registers a callback for SVC_Messages_svc_ServerSteamID
func (c *Callbacks) OnCSVCMsg_ServerSteamID(fn func(*dota.CSVCMsg_ServerSteamID) error) {
	c.onCSVCMsg_ServerSteamID = append(c.onCSVCMsg_ServerSteamID, fn)
}

// OnCSVCMsg_FullFrameSplit registers a callback for SVC_Messages_svc_FullFrameSplit
func (c *Callbacks) OnCSVCMsg_FullFrameSplit(fn func(*dota.CSVCMsg_FullFrameSplit) error) {
	c.onCSVCMsg_FullFrameSplit = append(c.onCSVCMsg_FullFrameSplit, fn)
}

// OnCSVCMsg_RconServerDetails registers a callback for SVC_Messages_svc_RconServerDetails
func (c *Callbacks) OnCSVCMsg_RconServerDetails(fn func(*dota.CSVCMsg_RconServerDetails) error) {
	c.onCSVCMsg_RconServerDetails = append(c.onCSVCMsg_RconServerDetails, fn)
}

// OnCSVCMsg_UserMessage registers a callback for SVC_Messages_svc_UserMessage
func (c *Callbacks) OnCSVCMsg_UserMessage(fn func(*dota.CSVCMsg_UserMessage) error) {
	c.onCSVCMsg_UserMessage = append(c.onCSVCMsg_UserMessage, fn)
}

// OnCSVCMsg_Broadcast_Command registers a callback for SVC_Messages_svc_Broadcast_Command
func (c *Callbacks) OnCSVCMsg_Broadcast_Command(fn func(*dota.CSVCMsg_Broadcast_Command) error) {
	c.onCSVCMsg_Broadcast_Command = append(c.onCSVCMsg_Broadcast_Command, fn)
}

// OnCSVCMsg_HltvFixupOperatorStatus registers a callback for SVC_Messages_svc_HltvFixupOperatorStatus
func (c *Callbacks) OnCSVCMsg_HltvFixupOperatorStatus(fn func(*dota.CSVCMsg_HltvFixupOperatorStatus) error) {
	c.onCSVCMsg_HltvFixupOperatorStatus = append(c.onCSVCMsg_HltvFixupOperatorStatus, fn)
}

// OnCUserMessageAchievementEvent registers a callback for EBaseUserMessages_UM_AchievementEvent
func (c *Callbacks) OnCUserMessageAchievementEvent(fn func(*dota.CUserMessageAchievementEvent) error) {
	c.onCUserMessageAchievementEvent = append(c.onCUserMessageAchievementEvent, fn)
}

// OnCUserMessageCloseCaption registers a callback for EBaseUserMessages_UM_CloseCaption
func (c *Callbacks) OnCUserMessageCloseCaption(fn func(*dota.CUserMessageCloseCaption) error) {
	c.onCUserMessageCloseCaption = append(c.onCUserMessageCloseCaption, fn)
}

// OnCUserMessageCloseCaptionDirect registers a callback for EBaseUserMessages_UM_CloseCaptionDirect
func (c *Callbacks) OnCUserMessageCloseCaptionDirect(fn func(*dota.CUserMessageCloseCaptionDirect) error) {
	c.onCUserMessageCloseCaptionDirect = append(c.onCUserMessageCloseCaptionDirect, fn)
}

// OnCUserMessageCurrentTimescale registers a callback for EBaseUserMessages_UM_CurrentTimescale
func (c *Callbacks) OnCUserMessageCurrentTimescale(fn func(*dota.CUserMessageCurrentTimescale) error) {
	c.onCUserMessageCurrentTimescale = append(c.onCUserMessageCurrentTimescale, fn)
}

// OnCUserMessageDesiredTimescale registers a callback for EBaseUserMessages_UM_DesiredTimescale
func (c *Callbacks) OnCUserMessageDesiredTimescale(fn func(*dota.CUserMessageDesiredTimescale) error) {
	c.onCUserMessageDesiredTimescale = append(c.onCUserMessageDesiredTimescale, fn)
}

// OnCUserMessageFade registers a callback for EBaseUserMessages_UM_Fade
func (c *Callbacks) OnCUserMessageFade(fn func(*dota.CUserMessageFade) error) {
	c.onCUserMessageFade = append(c.onCUserMessageFade, fn)
}

// OnCUserMessageGameTitle registers a callback for EBaseUserMessages_UM_GameTitle
func (c *Callbacks) OnCUserMessageGameTitle(fn func(*dota.CUserMessageGameTitle) error) {
	c.onCUserMessageGameTitle = append(c.onCUserMessageGameTitle, fn)
}

// OnCUserMessageHudMsg registers a callback for EBaseUserMessages_UM_HudMsg
func (c *Callbacks) OnCUserMessageHudMsg(fn func(*dota.CUserMessageHudMsg) error) {
	c.onCUserMessageHudMsg = append(c.onCUserMessageHudMsg, fn)
}

// OnCUserMessageHudText registers a callback for EBaseUserMessages_UM_HudText
func (c *Callbacks) OnCUserMessageHudText(fn func(*dota.CUserMessageHudText) error) {
	c.onCUserMessageHudText = append(c.onCUserMessageHudText, fn)
}

// OnCUserMessageColoredText registers a callback for EBaseUserMessages_UM_ColoredText
func (c *Callbacks) OnCUserMessageColoredText(fn func(*dota.CUserMessageColoredText) error) {
	c.onCUserMessageColoredText = append(c.onCUserMessageColoredText, fn)
}

// OnCUserMessageRequestState registers a callback for EBaseUserMessages_UM_RequestState
func (c *Callbacks) OnCUserMessageRequestState(fn func(*dota.CUserMessageRequestState) error) {
	c.onCUserMessageRequestState = append(c.onCUserMessageRequestState, fn)
}

// OnCUserMessageResetHUD registers a callback for EBaseUserMessages_UM_ResetHUD
func (c *Callbacks) OnCUserMessageResetHUD(fn func(*dota.CUserMessageResetHUD) error) {
	c.onCUserMessageResetHUD = append(c.onCUserMessageResetHUD, fn)
}

// OnCUserMessageRumble registers a callback for EBaseUserMessages_UM_Rumble
func (c *Callbacks) OnCUserMessageRumble(fn func(*dota.CUserMessageRumble) error) {
	c.onCUserMessageRumble = append(c.onCUserMessageRumble, fn)
}

// OnCUserMessageSayText registers a callback for EBaseUserMessages_UM_SayText
func (c *Callbacks) OnCUserMessageSayText(fn func(*dota.CUserMessageSayText) error) {
	c.onCUserMessageSayText = append(c.onCUserMessageSayText, fn)
}

// OnCUserMessageSayText2 registers a callback for EBaseUserMessages_UM_SayText2
func (c *Callbacks) OnCUserMessageSayText2(fn func(*dota.CUserMessageSayText2) error) {
	c.onCUserMessageSayText2 = append(c.onCUserMessageSayText2, fn)
}

// OnCUserMessageSayTextChannel registers a callback for EBaseUserMessages_UM_SayTextChannel
func (c *Callbacks) OnCUserMessageSayTextChannel(fn func(*dota.CUserMessageSayTextChannel) error) {
	c.onCUserMessageSayTextChannel = append(c.onCUserMessageSayTextChannel, fn)
}

// OnCUserMessageShake registers a callback for EBaseUserMessages_UM_Shake
func (c *Callbacks) OnCUserMessageShake(fn func(*dota.CUserMessageShake) error) {
	c.onCUserMessageShake = append(c.onCUserMessageShake, fn)
}

// OnCUserMessageShakeDir registers a callback for EBaseUserMessages_UM_ShakeDir
func (c *Callbacks) OnCUserMessageShakeDir(fn func(*dota.CUserMessageShakeDir) error) {
	c.onCUserMessageShakeDir = append(c.onCUserMessageShakeDir, fn)
}

// OnCUserMessageWaterShake registers a callback for EBaseUserMessages_UM_WaterShake
func (c *Callbacks) OnCUserMessageWaterShake(fn func(*dota.CUserMessageWaterShake) error) {
	c.onCUserMessageWaterShake = append(c.onCUserMessageWaterShake, fn)
}

// OnCUserMessageTextMsg registers a callback for EBaseUserMessages_UM_TextMsg
func (c *Callbacks) OnCUserMessageTextMsg(fn func(*dota.CUserMessageTextMsg) error) {
	c.onCUserMessageTextMsg = append(c.onCUserMessageTextMsg, fn)
}

// OnCUserMessageScreenTilt registers a callback for EBaseUserMessages_UM_ScreenTilt
func (c *Callbacks) OnCUserMessageScreenTilt(fn func(*dota.CUserMessageScreenTilt) error) {
	c.onCUserMessageScreenTilt = append(c.onCUserMessageScreenTilt, fn)
}

// OnCUserMessageVoiceMask registers a callback for EBaseUserMessages_UM_VoiceMask
func (c *Callbacks) OnCUserMessageVoiceMask(fn func(*dota.CUserMessageVoiceMask) error) {
	c.onCUserMessageVoiceMask = append(c.onCUserMessageVoiceMask, fn)
}

// OnCUserMessageSendAudio registers a callback for EBaseUserMessages_UM_SendAudio
func (c *Callbacks) OnCUserMessageSendAudio(fn func(*dota.CUserMessageSendAudio) error) {
	c.onCUserMessageSendAudio = append(c.onCUserMessageSendAudio, fn)
}

// OnCUserMessageItemPickup registers a callback for EBaseUserMessages_UM_ItemPickup
func (c *Callbacks) OnCUserMessageItemPickup(fn func(*dota.CUserMessageItemPickup) error) {
	c.onCUserMessageItemPickup = append(c.onCUserMessageItemPickup, fn)
}

// OnCUserMessageAmmoDenied registers a callback for EBaseUserMessages_UM_AmmoDenied
func (c *Callbacks) OnCUserMessageAmmoDenied(fn func(*dota.CUserMessageAmmoDenied) error) {
	c.onCUserMessageAmmoDenied = append(c.onCUserMessageAmmoDenied, fn)
}

// OnCUserMessageShowMenu registers a callback for EBaseUserMessages_UM_ShowMenu
func (c *Callbacks) OnCUserMessageShowMenu(fn func(*dota.CUserMessageShowMenu) error) {
	c.onCUserMessageShowMenu = append(c.onCUserMessageShowMenu, fn)
}

// OnCUserMessageCreditsMsg registers a callback for EBaseUserMessages_UM_CreditsMsg
func (c *Callbacks) OnCUserMessageCreditsMsg(fn func(*dota.CUserMessageCreditsMsg) error) {
	c.onCUserMessageCreditsMsg = append(c.onCUserMessageCreditsMsg, fn)
}

// OnCEntityMessagePlayJingle registers a callback for EBaseEntityMessages_EM_PlayJingle
func (c *Callbacks) OnCEntityMessagePlayJingle(fn func(*dota.CEntityMessagePlayJingle) error) {
	c.onCEntityMessagePlayJingle = append(c.onCEntityMessagePlayJingle, fn)
}

// OnCEntityMessageScreenOverlay registers a callback for EBaseEntityMessages_EM_ScreenOverlay
func (c *Callbacks) OnCEntityMessageScreenOverlay(fn func(*dota.CEntityMessageScreenOverlay) error) {
	c.onCEntityMessageScreenOverlay = append(c.onCEntityMessageScreenOverlay, fn)
}

// OnCEntityMessageRemoveAllDecals registers a callback for EBaseEntityMessages_EM_RemoveAllDecals
func (c *Callbacks) OnCEntityMessageRemoveAllDecals(fn func(*dota.CEntityMessageRemoveAllDecals) error) {
	c.onCEntityMessageRemoveAllDecals = append(c.onCEntityMessageRemoveAllDecals, fn)
}

// OnCEntityMessagePropagateForce registers a callback for EBaseEntityMessages_EM_PropagateForce
func (c *Callbacks) OnCEntityMessagePropagateForce(fn func(*dota.CEntityMessagePropagateForce) error) {
	c.onCEntityMessagePropagateForce = append(c.onCEntityMessagePropagateForce, fn)
}

// OnCEntityMessageDoSpark registers a callback for EBaseEntityMessages_EM_DoSpark
func (c *Callbacks) OnCEntityMessageDoSpark(fn func(*dota.CEntityMessageDoSpark) error) {
	c.onCEntityMessageDoSpark = append(c.onCEntityMessageDoSpark, fn)
}

// OnCEntityMessageFixAngle registers a callback for EBaseEntityMessages_EM_FixAngle
func (c *Callbacks) OnCEntityMessageFixAngle(fn func(*dota.CEntityMessageFixAngle) error) {
	c.onCEntityMessageFixAngle = append(c.onCEntityMessageFixAngle, fn)
}

// OnCUserMessageCloseCaptionPlaceholder registers a callback for EBaseUserMessages_UM_CloseCaptionPlaceholder
func (c *Callbacks) OnCUserMessageCloseCaptionPlaceholder(fn func(*dota.CUserMessageCloseCaptionPlaceholder) error) {
	c.onCUserMessageCloseCaptionPlaceholder = append(c.onCUserMessageCloseCaptionPlaceholder, fn)
}

// OnCUserMessageCameraTransition registers a callback for EBaseUserMessages_UM_CameraTransition
func (c *Callbacks) OnCUserMessageCameraTransition(fn func(*dota.CUserMessageCameraTransition) error) {
	c.onCUserMessageCameraTransition = append(c.onCUserMessageCameraTransition, fn)
}

// OnCUserMessageAudioParameter registers a callback for EBaseUserMessages_UM_AudioParameter
func (c *Callbacks) OnCUserMessageAudioParameter(fn func(*dota.CUserMessageAudioParameter) error) {
	c.onCUserMessageAudioParameter = append(c.onCUserMessageAudioParameter, fn)
}

// OnCUserMessageHapticsManagerPulse registers a callback for EBaseUserMessages_UM_HapticsManagerPulse
func (c *Callbacks) OnCUserMessageHapticsManagerPulse(fn func(*dota.CUserMessageHapticsManagerPulse) error) {
	c.onCUserMessageHapticsManagerPulse = append(c.onCUserMessageHapticsManagerPulse, fn)
}

// OnCUserMessageHapticsManagerEffect registers a callback for EBaseUserMessages_UM_HapticsManagerEffect
func (c *Callbacks) OnCUserMessageHapticsManagerEffect(fn func(*dota.CUserMessageHapticsManagerEffect) error) {
	c.onCUserMessageHapticsManagerEffect = append(c.onCUserMessageHapticsManagerEffect, fn)
}

// OnCUserMessageUpdateCssClasses registers a callback for EBaseUserMessages_UM_UpdateCssClasses
func (c *Callbacks) OnCUserMessageUpdateCssClasses(fn func(*dota.CUserMessageUpdateCssClasses) error) {
	c.onCUserMessageUpdateCssClasses = append(c.onCUserMessageUpdateCssClasses, fn)
}

// OnCUserMessageServerFrameTime registers a callback for EBaseUserMessages_UM_ServerFrameTime
func (c *Callbacks) OnCUserMessageServerFrameTime(fn func(*dota.CUserMessageServerFrameTime) error) {
	c.onCUserMessageServerFrameTime = append(c.onCUserMessageServerFrameTime, fn)
}

// OnCUserMessageLagCompensationError registers a callback for EBaseUserMessages_UM_LagCompensationError
func (c *Callbacks) OnCUserMessageLagCompensationError(fn func(*dota.CUserMessageLagCompensationError) error) {
	c.onCUserMessageLagCompensationError = append(c.onCUserMessageLagCompensationError, fn)
}

// OnCUserMessageRequestDllStatus registers a callback for EBaseUserMessages_UM_RequestDllStatus
func (c *Callbacks) OnCUserMessageRequestDllStatus(fn func(*dota.CUserMessageRequestDllStatus) error) {
	c.onCUserMessageRequestDllStatus = append(c.onCUserMessageRequestDllStatus, fn)
}

// OnCUserMessageRequestUtilAction registers a callback for EBaseUserMessages_UM_RequestUtilAction
func (c *Callbacks) OnCUserMessageRequestUtilAction(fn func(*dota.CUserMessageRequestUtilAction) error) {
	c.onCUserMessageRequestUtilAction = append(c.onCUserMessageRequestUtilAction, fn)
}

// OnCUserMessageRequestInventory registers a callback for EBaseUserMessages_UM_RequestInventory
func (c *Callbacks) OnCUserMessageRequestInventory(fn func(*dota.CUserMessageRequestInventory) error) {
	c.onCUserMessageRequestInventory = append(c.onCUserMessageRequestInventory, fn)
}

// OnCUserMessageRequestDiagnostic registers a callback for EBaseUserMessages_UM_RequestDiagnostic
func (c *Callbacks) OnCUserMessageRequestDiagnostic(fn func(*dota.CUserMessageRequestDiagnostic) error) {
	c.onCUserMessageRequestDiagnostic = append(c.onCUserMessageRequestDiagnostic, fn)
}

// OnCMsgVDebugGameSessionIDEvent registers a callback for EBaseGameEvents_GE_VDebugGameSessionIDEvent
func (c *Callbacks) OnCMsgVDebugGameSessionIDEvent(fn func(*dota.CMsgVDebugGameSessionIDEvent) error) {
	c.onCMsgVDebugGameSessionIDEvent = append(c.onCMsgVDebugGameSessionIDEvent, fn)
}

// OnCMsgPlaceDecalEvent registers a callback for EBaseGameEvents_GE_PlaceDecalEvent
func (c *Callbacks) OnCMsgPlaceDecalEvent(fn func(*dota.CMsgPlaceDecalEvent) error) {
	c.onCMsgPlaceDecalEvent = append(c.onCMsgPlaceDecalEvent, fn)
}

// OnCMsgClearWorldDecalsEvent registers a callback for EBaseGameEvents_GE_ClearWorldDecalsEvent
func (c *Callbacks) OnCMsgClearWorldDecalsEvent(fn func(*dota.CMsgClearWorldDecalsEvent) error) {
	c.onCMsgClearWorldDecalsEvent = append(c.onCMsgClearWorldDecalsEvent, fn)
}

// OnCMsgClearEntityDecalsEvent registers a callback for EBaseGameEvents_GE_ClearEntityDecalsEvent
func (c *Callbacks) OnCMsgClearEntityDecalsEvent(fn func(*dota.CMsgClearEntityDecalsEvent) error) {
	c.onCMsgClearEntityDecalsEvent = append(c.onCMsgClearEntityDecalsEvent, fn)
}

// OnCMsgClearDecalsForSkeletonInstanceEvent registers a callback for EBaseGameEvents_GE_ClearDecalsForSkeletonInstanceEvent
func (c *Callbacks) OnCMsgClearDecalsForSkeletonInstanceEvent(fn func(*dota.CMsgClearDecalsForSkeletonInstanceEvent) error) {
	c.onCMsgClearDecalsForSkeletonInstanceEvent = append(c.onCMsgClearDecalsForSkeletonInstanceEvent, fn)
}

// OnCMsgSource1LegacyGameEventList registers a callback for EBaseGameEvents_GE_Source1LegacyGameEventList
func (c *Callbacks) OnCMsgSource1LegacyGameEventList(fn func(*dota.CMsgSource1LegacyGameEventList) error) {
	c.onCMsgSource1LegacyGameEventList = append(c.onCMsgSource1LegacyGameEventList, fn)
}

// OnCMsgSource1LegacyListenEvents registers a callback for EBaseGameEvents_GE_Source1LegacyListenEvents
func (c *Callbacks) OnCMsgSource1LegacyListenEvents(fn func(*dota.CMsgSource1LegacyListenEvents) error) {
	c.onCMsgSource1LegacyListenEvents = append(c.onCMsgSource1LegacyListenEvents, fn)
}

// OnCMsgSource1LegacyGameEvent registers a callback for EBaseGameEvents_GE_Source1LegacyGameEvent
func (c *Callbacks) OnCMsgSource1LegacyGameEvent(fn func(*dota.CMsgSource1LegacyGameEvent) error) {
	c.onCMsgSource1LegacyGameEvent = append(c.onCMsgSource1LegacyGameEvent, fn)
}

// OnCMsgSosStartSoundEvent registers a callback for EBaseGameEvents_GE_SosStartSoundEvent
func (c *Callbacks) OnCMsgSosStartSoundEvent(fn func(*dota.CMsgSosStartSoundEvent) error) {
	c.onCMsgSosStartSoundEvent = append(c.onCMsgSosStartSoundEvent, fn)
}

// OnCMsgSosStopSoundEvent registers a callback for EBaseGameEvents_GE_SosStopSoundEvent
func (c *Callbacks) OnCMsgSosStopSoundEvent(fn func(*dota.CMsgSosStopSoundEvent) error) {
	c.onCMsgSosStopSoundEvent = append(c.onCMsgSosStopSoundEvent, fn)
}

// OnCMsgSosSetSoundEventParams registers a callback for EBaseGameEvents_GE_SosSetSoundEventParams
func (c *Callbacks) OnCMsgSosSetSoundEventParams(fn func(*dota.CMsgSosSetSoundEventParams) error) {
	c.onCMsgSosSetSoundEventParams = append(c.onCMsgSosSetSoundEventParams, fn)
}

// OnCMsgSosSetLibraryStackFields registers a callback for EBaseGameEvents_GE_SosSetLibraryStackFields
func (c *Callbacks) OnCMsgSosSetLibraryStackFields(fn func(*dota.CMsgSosSetLibraryStackFields) error) {
	c.onCMsgSosSetLibraryStackFields = append(c.onCMsgSosSetLibraryStackFields, fn)
}

// OnCMsgSosStopSoundEventHash registers a callback for EBaseGameEvents_GE_SosStopSoundEventHash
func (c *Callbacks) OnCMsgSosStopSoundEventHash(fn func(*dota.CMsgSosStopSoundEventHash) error) {
	c.onCMsgSosStopSoundEventHash = append(c.onCMsgSosStopSoundEventHash, fn)
}

// OnCCitadelUserMessage_Damage registers a callback for CitadelUserMessageIds_k_EUserMsg_Damage
func (c *Callbacks) OnCCitadelUserMessage_Damage(fn func(*dota.CCitadelUserMessage_Damage) error) {
	c.onCCitadelUserMessage_Damage = append(c.onCCitadelUserMessage_Damage, fn)
}

// OnCCitadelUserMsg_MapPing registers a callback for CitadelUserMessageIds_k_EUserMsg_MapPing
func (c *Callbacks) OnCCitadelUserMsg_MapPing(fn func(*dota.CCitadelUserMsg_MapPing) error) {
	c.onCCitadelUserMsg_MapPing = append(c.onCCitadelUserMsg_MapPing, fn)
}

// OnCCitadelUserMsg_TeamRewards registers a callback for CitadelUserMessageIds_k_EUserMsg_TeamRewards
func (c *Callbacks) OnCCitadelUserMsg_TeamRewards(fn func(*dota.CCitadelUserMsg_TeamRewards) error) {
	c.onCCitadelUserMsg_TeamRewards = append(c.onCCitadelUserMsg_TeamRewards, fn)
}

// OnCCitadelUserMsg_TriggerDamageFlash registers a callback for CitadelUserMessageIds_k_EUserMsg_TriggerDamageFlash
func (c *Callbacks) OnCCitadelUserMsg_TriggerDamageFlash(fn func(*dota.CCitadelUserMsg_TriggerDamageFlash) error) {
	c.onCCitadelUserMsg_TriggerDamageFlash = append(c.onCCitadelUserMsg_TriggerDamageFlash, fn)
}

// OnCCitadelUserMsg_AbilitiesChanged registers a callback for CitadelUserMessageIds_k_EUserMsg_AbilitiesChanged
func (c *Callbacks) OnCCitadelUserMsg_AbilitiesChanged(fn func(*dota.CCitadelUserMsg_AbilitiesChanged) error) {
	c.onCCitadelUserMsg_AbilitiesChanged = append(c.onCCitadelUserMsg_AbilitiesChanged, fn)
}

// OnCCitadelUserMsg_RecentDamageSummary registers a callback for CitadelUserMessageIds_k_EUserMsg_RecentDamageSummary
func (c *Callbacks) OnCCitadelUserMsg_RecentDamageSummary(fn func(*dota.CCitadelUserMsg_RecentDamageSummary) error) {
	c.onCCitadelUserMsg_RecentDamageSummary = append(c.onCCitadelUserMsg_RecentDamageSummary, fn)
}

// OnCCitadelUserMsg_SpectatorTeamChanged registers a callback for CitadelUserMessageIds_k_EUserMsg_SpectatorTeamChanged
func (c *Callbacks) OnCCitadelUserMsg_SpectatorTeamChanged(fn func(*dota.CCitadelUserMsg_SpectatorTeamChanged) error) {
	c.onCCitadelUserMsg_SpectatorTeamChanged = append(c.onCCitadelUserMsg_SpectatorTeamChanged, fn)
}

// OnCCitadelUserMsg_ChatWheel registers a callback for CitadelUserMessageIds_k_EUserMsg_ChatWheel
func (c *Callbacks) OnCCitadelUserMsg_ChatWheel(fn func(*dota.CCitadelUserMsg_ChatWheel) error) {
	c.onCCitadelUserMsg_ChatWheel = append(c.onCCitadelUserMsg_ChatWheel, fn)
}

// OnCCitadelUserMsg_GoldHistory registers a callback for CitadelUserMessageIds_k_EUserMsg_GoldHistory
func (c *Callbacks) OnCCitadelUserMsg_GoldHistory(fn func(*dota.CCitadelUserMsg_GoldHistory) error) {
	c.onCCitadelUserMsg_GoldHistory = append(c.onCCitadelUserMsg_GoldHistory, fn)
}

// OnCCitadelUserMsg_ChatMsg registers a callback for CitadelUserMessageIds_k_EUserMsg_ChatMsg
func (c *Callbacks) OnCCitadelUserMsg_ChatMsg(fn func(*dota.CCitadelUserMsg_ChatMsg) error) {
	c.onCCitadelUserMsg_ChatMsg = append(c.onCCitadelUserMsg_ChatMsg, fn)
}

// OnCCitadelUserMsg_QuickResponse registers a callback for CitadelUserMessageIds_k_EUserMsg_QuickResponse
func (c *Callbacks) OnCCitadelUserMsg_QuickResponse(fn func(*dota.CCitadelUserMsg_QuickResponse) error) {
	c.onCCitadelUserMsg_QuickResponse = append(c.onCCitadelUserMsg_QuickResponse, fn)
}

// OnCCitadelUserMsg_PostMatchDetails registers a callback for CitadelUserMessageIds_k_EUserMsg_PostMatchDetails
func (c *Callbacks) OnCCitadelUserMsg_PostMatchDetails(fn func(*dota.CCitadelUserMsg_PostMatchDetails) error) {
	c.onCCitadelUserMsg_PostMatchDetails = append(c.onCCitadelUserMsg_PostMatchDetails, fn)
}

// OnCCitadelUserMsg_ChatEvent registers a callback for CitadelUserMessageIds_k_EUserMsg_ChatEvent
func (c *Callbacks) OnCCitadelUserMsg_ChatEvent(fn func(*dota.CCitadelUserMsg_ChatEvent) error) {
	c.onCCitadelUserMsg_ChatEvent = append(c.onCCitadelUserMsg_ChatEvent, fn)
}

// OnCCitadelUserMsg_AbilityInterrupted registers a callback for CitadelUserMessageIds_k_EUserMsg_AbilityInterrupted
func (c *Callbacks) OnCCitadelUserMsg_AbilityInterrupted(fn func(*dota.CCitadelUserMsg_AbilityInterrupted) error) {
	c.onCCitadelUserMsg_AbilityInterrupted = append(c.onCCitadelUserMsg_AbilityInterrupted, fn)
}

// OnCCitadelUserMsg_HeroKilled registers a callback for CitadelUserMessageIds_k_EUserMsg_HeroKilled
func (c *Callbacks) OnCCitadelUserMsg_HeroKilled(fn func(*dota.CCitadelUserMsg_HeroKilled) error) {
	c.onCCitadelUserMsg_HeroKilled = append(c.onCCitadelUserMsg_HeroKilled, fn)
}

// OnCCitadelUserMsg_ReturnIdol registers a callback for CitadelUserMessageIds_k_EUserMsg_ReturnIdol
func (c *Callbacks) OnCCitadelUserMsg_ReturnIdol(fn func(*dota.CCitadelUserMsg_ReturnIdol) error) {
	c.onCCitadelUserMsg_ReturnIdol = append(c.onCCitadelUserMsg_ReturnIdol, fn)
}

// OnCCitadelUserMsg_SetClientCameraAngles registers a callback for CitadelUserMessageIds_k_EUserMsg_SetClientCameraAngles
func (c *Callbacks) OnCCitadelUserMsg_SetClientCameraAngles(fn func(*dota.CCitadelUserMsg_SetClientCameraAngles) error) {
	c.onCCitadelUserMsg_SetClientCameraAngles = append(c.onCCitadelUserMsg_SetClientCameraAngles, fn)
}

// OnCCitadelUserMsg_MapLine registers a callback for CitadelUserMessageIds_k_EUserMsg_MapLine
func (c *Callbacks) OnCCitadelUserMsg_MapLine(fn func(*dota.CCitadelUserMsg_MapLine) error) {
	c.onCCitadelUserMsg_MapLine = append(c.onCCitadelUserMsg_MapLine, fn)
}

// OnCCitadelUserMessage_BulletHit registers a callback for CitadelUserMessageIds_k_EUserMsg_BulletHit
func (c *Callbacks) OnCCitadelUserMessage_BulletHit(fn func(*dota.CCitadelUserMessage_BulletHit) error) {
	c.onCCitadelUserMessage_BulletHit = append(c.onCCitadelUserMessage_BulletHit, fn)
}

// OnCCitadelUserMessage_ObjectiveMask registers a callback for CitadelUserMessageIds_k_EUserMsg_ObjectiveMask
func (c *Callbacks) OnCCitadelUserMessage_ObjectiveMask(fn func(*dota.CCitadelUserMessage_ObjectiveMask) error) {
	c.onCCitadelUserMessage_ObjectiveMask = append(c.onCCitadelUserMessage_ObjectiveMask, fn)
}

// OnCCitadelUserMessage_ModifierApplied registers a callback for CitadelUserMessageIds_k_EUserMsg_ModifierApplied
func (c *Callbacks) OnCCitadelUserMessage_ModifierApplied(fn func(*dota.CCitadelUserMessage_ModifierApplied) error) {
	c.onCCitadelUserMessage_ModifierApplied = append(c.onCCitadelUserMessage_ModifierApplied, fn)
}

// OnCCitadelUserMsg_CameraController registers a callback for CitadelUserMessageIds_k_EUserMsg_CameraController
func (c *Callbacks) OnCCitadelUserMsg_CameraController(fn func(*dota.CCitadelUserMsg_CameraController) error) {
	c.onCCitadelUserMsg_CameraController = append(c.onCCitadelUserMsg_CameraController, fn)
}

// OnCCitadelUserMessage_AuraModifierApplied registers a callback for CitadelUserMessageIds_k_EUserMsg_AuraModifierApplied
func (c *Callbacks) OnCCitadelUserMessage_AuraModifierApplied(fn func(*dota.CCitadelUserMessage_AuraModifierApplied) error) {
	c.onCCitadelUserMessage_AuraModifierApplied = append(c.onCCitadelUserMessage_AuraModifierApplied, fn)
}

// OnCCitadelUserMsg_ObstructedShotFired registers a callback for CitadelUserMessageIds_k_EUserMsg_ObstructedShotFired
func (c *Callbacks) OnCCitadelUserMsg_ObstructedShotFired(fn func(*dota.CCitadelUserMsg_ObstructedShotFired) error) {
	c.onCCitadelUserMsg_ObstructedShotFired = append(c.onCCitadelUserMsg_ObstructedShotFired, fn)
}

// OnCCitadelUserMsg_AbilityLateFailure registers a callback for CitadelUserMessageIds_k_EUserMsg_AbilityLateFailure
func (c *Callbacks) OnCCitadelUserMsg_AbilityLateFailure(fn func(*dota.CCitadelUserMsg_AbilityLateFailure) error) {
	c.onCCitadelUserMsg_AbilityLateFailure = append(c.onCCitadelUserMsg_AbilityLateFailure, fn)
}

// OnCCitadelUserMsg_AbilityPing registers a callback for CitadelUserMessageIds_k_EUserMsg_AbilityPing
func (c *Callbacks) OnCCitadelUserMsg_AbilityPing(fn func(*dota.CCitadelUserMsg_AbilityPing) error) {
	c.onCCitadelUserMsg_AbilityPing = append(c.onCCitadelUserMsg_AbilityPing, fn)
}

// OnCCitadelUserMsg_PostProcessingAnim registers a callback for CitadelUserMessageIds_k_EUserMsg_PostProcessingAnim
func (c *Callbacks) OnCCitadelUserMsg_PostProcessingAnim(fn func(*dota.CCitadelUserMsg_PostProcessingAnim) error) {
	c.onCCitadelUserMsg_PostProcessingAnim = append(c.onCCitadelUserMsg_PostProcessingAnim, fn)
}

// OnCCitadelUserMsg_DeathReplayData registers a callback for CitadelUserMessageIds_k_EUserMsg_DeathReplayData
func (c *Callbacks) OnCCitadelUserMsg_DeathReplayData(fn func(*dota.CCitadelUserMsg_DeathReplayData) error) {
	c.onCCitadelUserMsg_DeathReplayData = append(c.onCCitadelUserMsg_DeathReplayData, fn)
}

// OnCCitadelUserMsg_PlayerLifetimeStatInfo registers a callback for CitadelUserMessageIds_k_EUserMsg_PlayerLifetimeStatInfo
func (c *Callbacks) OnCCitadelUserMsg_PlayerLifetimeStatInfo(fn func(*dota.CCitadelUserMsg_PlayerLifetimeStatInfo) error) {
	c.onCCitadelUserMsg_PlayerLifetimeStatInfo = append(c.onCCitadelUserMsg_PlayerLifetimeStatInfo, fn)
}

// OnCCitadelUserMsg_ForceShopClosed registers a callback for CitadelUserMessageIds_k_EUserMsg_ForceShopClosed
func (c *Callbacks) OnCCitadelUserMsg_ForceShopClosed(fn func(*dota.CCitadelUserMsg_ForceShopClosed) error) {
	c.onCCitadelUserMsg_ForceShopClosed = append(c.onCCitadelUserMsg_ForceShopClosed, fn)
}

// OnCCitadelUserMsg_StaminaDrained registers a callback for CitadelUserMessageIds_k_EUserMsg_StaminaDrained
func (c *Callbacks) OnCCitadelUserMsg_StaminaDrained(fn func(*dota.CCitadelUserMsg_StaminaDrained) error) {
	c.onCCitadelUserMsg_StaminaDrained = append(c.onCCitadelUserMsg_StaminaDrained, fn)
}

// OnCCitadelUserMessage_AbilityNotify registers a callback for CitadelUserMessageIds_k_EUserMsg_AbilityNotify
func (c *Callbacks) OnCCitadelUserMessage_AbilityNotify(fn func(*dota.CCitadelUserMessage_AbilityNotify) error) {
	c.onCCitadelUserMessage_AbilityNotify = append(c.onCCitadelUserMessage_AbilityNotify, fn)
}

// OnCCitadelUserMsg_GetDamageStatsResponse registers a callback for CitadelUserMessageIds_k_EUserMsg_GetDamageStatsResponse
func (c *Callbacks) OnCCitadelUserMsg_GetDamageStatsResponse(fn func(*dota.CCitadelUserMsg_GetDamageStatsResponse) error) {
	c.onCCitadelUserMsg_GetDamageStatsResponse = append(c.onCCitadelUserMsg_GetDamageStatsResponse, fn)
}

// OnCCitadelUserMsg_ParticipantStartSoundEvent registers a callback for CitadelUserMessageIds_k_EUserMsg_ParticipantStartSoundEvent
func (c *Callbacks) OnCCitadelUserMsg_ParticipantStartSoundEvent(fn func(*dota.CCitadelUserMsg_ParticipantStartSoundEvent) error) {
	c.onCCitadelUserMsg_ParticipantStartSoundEvent = append(c.onCCitadelUserMsg_ParticipantStartSoundEvent, fn)
}

// OnCCitadelUserMsg_ParticipantStopSoundEvent registers a callback for CitadelUserMessageIds_k_EUserMsg_ParticipantStopSoundEvent
func (c *Callbacks) OnCCitadelUserMsg_ParticipantStopSoundEvent(fn func(*dota.CCitadelUserMsg_ParticipantStopSoundEvent) error) {
	c.onCCitadelUserMsg_ParticipantStopSoundEvent = append(c.onCCitadelUserMsg_ParticipantStopSoundEvent, fn)
}

// OnCCitadelUserMsg_ParticipantStopSoundEventHash registers a callback for CitadelUserMessageIds_k_EUserMsg_ParticipantStopSoundEventHash
func (c *Callbacks) OnCCitadelUserMsg_ParticipantStopSoundEventHash(fn func(*dota.CCitadelUserMsg_ParticipantStopSoundEventHash) error) {
	c.onCCitadelUserMsg_ParticipantStopSoundEventHash = append(c.onCCitadelUserMsg_ParticipantStopSoundEventHash, fn)
}

// OnCCitadelUserMsg_ParticipantSetSoundEventParams registers a callback for CitadelUserMessageIds_k_EUserMsg_ParticipantSetSoundEventParams
func (c *Callbacks) OnCCitadelUserMsg_ParticipantSetSoundEventParams(fn func(*dota.CCitadelUserMsg_ParticipantSetSoundEventParams) error) {
	c.onCCitadelUserMsg_ParticipantSetSoundEventParams = append(c.onCCitadelUserMsg_ParticipantSetSoundEventParams, fn)
}

// OnCCitadelUserMsg_ParticipantSetLibraryStackFields registers a callback for CitadelUserMessageIds_k_EUserMsg_ParticipantSetLibraryStackFields
func (c *Callbacks) OnCCitadelUserMsg_ParticipantSetLibraryStackFields(fn func(*dota.CCitadelUserMsg_ParticipantSetLibraryStackFields) error) {
	c.onCCitadelUserMsg_ParticipantSetLibraryStackFields = append(c.onCCitadelUserMsg_ParticipantSetLibraryStackFields, fn)
}

// OnCCitadelUserMsg_BossKilled registers a callback for CitadelUserMessageIds_k_EUserMsg_BossKilled
func (c *Callbacks) OnCCitadelUserMsg_BossKilled(fn func(*dota.CCitadelUserMsg_BossKilled) error) {
	c.onCCitadelUserMsg_BossKilled = append(c.onCCitadelUserMsg_BossKilled, fn)
}

// OnCMsgFireBullets registers a callback for ECitadelGameEvents_GE_FireBullets
func (c *Callbacks) OnCMsgFireBullets(fn func(*dota.CMsgFireBullets) error) {
	c.onCMsgFireBullets = append(c.onCMsgFireBullets, fn)
}

// OnCMsgPlayerAnimEvent registers a callback for ECitadelGameEvents_GE_PlayerAnimEvent
func (c *Callbacks) OnCMsgPlayerAnimEvent(fn func(*dota.CMsgPlayerAnimEvent) error) {
	c.onCMsgPlayerAnimEvent = append(c.onCMsgPlayerAnimEvent, fn)
}

// OnCMsgParticleSystemManager registers a callback for ECitadelGameEvents_GE_ParticleSystemManager
func (c *Callbacks) OnCMsgParticleSystemManager(fn func(*dota.CMsgParticleSystemManager) error) {
	c.onCMsgParticleSystemManager = append(c.onCMsgParticleSystemManager, fn)
}

// OnCMsgScreenTextPretty registers a callback for ECitadelGameEvents_GE_ScreenTextPretty
func (c *Callbacks) OnCMsgScreenTextPretty(fn func(*dota.CMsgScreenTextPretty) error) {
	c.onCMsgScreenTextPretty = append(c.onCMsgScreenTextPretty, fn)
}

// OnCMsgServerRequestedTracer registers a callback for ECitadelGameEvents_GE_ServerRequestedTracer
func (c *Callbacks) OnCMsgServerRequestedTracer(fn func(*dota.CMsgServerRequestedTracer) error) {
	c.onCMsgServerRequestedTracer = append(c.onCMsgServerRequestedTracer, fn)
}

// OnCMsgBulletImpact registers a callback for ECitadelGameEvents_GE_BulletImpact
func (c *Callbacks) OnCMsgBulletImpact(fn func(*dota.CMsgBulletImpact) error) {
	c.onCMsgBulletImpact = append(c.onCMsgBulletImpact, fn)
}

// OnCMsgEnableSatVolumesEvent registers a callback for ECitadelGameEvents_GE_EnableSatVolumesEvent
func (c *Callbacks) OnCMsgEnableSatVolumesEvent(fn func(*dota.CMsgEnableSatVolumesEvent) error) {
	c.onCMsgEnableSatVolumesEvent = append(c.onCMsgEnableSatVolumesEvent, fn)
}

// OnCMsgPlaceSatVolumeEvent registers a callback for ECitadelGameEvents_GE_PlaceSatVolumeEvent
func (c *Callbacks) OnCMsgPlaceSatVolumeEvent(fn func(*dota.CMsgPlaceSatVolumeEvent) error) {
	c.onCMsgPlaceSatVolumeEvent = append(c.onCMsgPlaceSatVolumeEvent, fn)
}

// OnCMsgDisableSatVolumesEvent registers a callback for ECitadelGameEvents_GE_DisableSatVolumesEvent
func (c *Callbacks) OnCMsgDisableSatVolumesEvent(fn func(*dota.CMsgDisableSatVolumesEvent) error) {
	c.onCMsgDisableSatVolumesEvent = append(c.onCMsgDisableSatVolumesEvent, fn)
}

// OnCMsgRemoveSatVolumeEvent registers a callback for ECitadelGameEvents_GE_RemoveSatVolumeEvent
func (c *Callbacks) OnCMsgRemoveSatVolumeEvent(fn func(*dota.CMsgRemoveSatVolumeEvent) error) {
	c.onCMsgRemoveSatVolumeEvent = append(c.onCMsgRemoveSatVolumeEvent, fn)
}

func (c *Callbacks) callByDemoType(t int32, buf []byte) error {
	switch t {
	case 0: // dota.EDemoCommands_DEM_Stop
		if c.onCDemoStop == nil {
			return nil
		}

		msg := &dota.CDemoStop{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoStop {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 1: // dota.EDemoCommands_DEM_FileHeader
		if c.onCDemoFileHeader == nil {
			return nil
		}

		msg := &dota.CDemoFileHeader{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoFileHeader {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 2: // dota.EDemoCommands_DEM_FileInfo
		if c.onCDemoFileInfo == nil {
			return nil
		}

		msg := &dota.CDemoFileInfo{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoFileInfo {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 3: // dota.EDemoCommands_DEM_SyncTick
		if c.onCDemoSyncTick == nil {
			return nil
		}

		msg := &dota.CDemoSyncTick{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoSyncTick {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 4: // dota.EDemoCommands_DEM_SendTables
		if c.onCDemoSendTables == nil {
			return nil
		}

		msg := &dota.CDemoSendTables{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoSendTables {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 5: // dota.EDemoCommands_DEM_ClassInfo
		if c.onCDemoClassInfo == nil {
			return nil
		}

		msg := &dota.CDemoClassInfo{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoClassInfo {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 6: // dota.EDemoCommands_DEM_StringTables
		if c.onCDemoStringTables == nil {
			return nil
		}

		msg := &dota.CDemoStringTables{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoStringTables {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 7: // dota.EDemoCommands_DEM_Packet
		if c.onCDemoPacket == nil {
			return nil
		}

		msg := &dota.CDemoPacket{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoPacket {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 8: // dota.EDemoCommands_DEM_SignonPacket
		if c.onCDemoSignonPacket == nil {
			return nil
		}

		msg := &dota.CDemoPacket{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoSignonPacket {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 9: // dota.EDemoCommands_DEM_ConsoleCmd
		if c.onCDemoConsoleCmd == nil {
			return nil
		}

		msg := &dota.CDemoConsoleCmd{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoConsoleCmd {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 10: // dota.EDemoCommands_DEM_CustomData
		if c.onCDemoCustomData == nil {
			return nil
		}

		msg := &dota.CDemoCustomData{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoCustomData {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 11: // dota.EDemoCommands_DEM_CustomDataCallbacks
		if c.onCDemoCustomDataCallbacks == nil {
			return nil
		}

		msg := &dota.CDemoCustomDataCallbacks{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoCustomDataCallbacks {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 12: // dota.EDemoCommands_DEM_UserCmd
		if c.onCDemoUserCmd == nil {
			return nil
		}

		msg := &dota.CDemoUserCmd{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoUserCmd {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 13: // dota.EDemoCommands_DEM_FullPacket
		if c.onCDemoFullPacket == nil {
			return nil
		}

		msg := &dota.CDemoFullPacket{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoFullPacket {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 14: // dota.EDemoCommands_DEM_SaveGame
		if c.onCDemoSaveGame == nil {
			return nil
		}

		msg := &dota.CDemoSaveGame{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoSaveGame {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 15: // dota.EDemoCommands_DEM_SpawnGroups
		if c.onCDemoSpawnGroups == nil {
			return nil
		}

		msg := &dota.CDemoSpawnGroups{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoSpawnGroups {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 16: // dota.EDemoCommands_DEM_AnimationData
		if c.onCDemoAnimationData == nil {
			return nil
		}

		msg := &dota.CDemoAnimationData{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoAnimationData {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 17: // dota.EDemoCommands_DEM_AnimationHeader
		if c.onCDemoAnimationHeader == nil {
			return nil
		}

		msg := &dota.CDemoAnimationHeader{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCDemoAnimationHeader {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	}

	if v(1) {
		_debugf("warning: no demo type %d found", t)
	}

	return nil
}

func (c *Callbacks) callByPacketType(t int32, buf []byte) error {
	switch t {
	case 0: // dota.NET_Messages_net_NOP
		if c.onCNETMsg_NOP == nil {
			return nil
		}

		msg := &dota.CNETMsg_NOP{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_NOP {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 3: // dota.NET_Messages_net_SplitScreenUser
		if c.onCNETMsg_SplitScreenUser == nil {
			return nil
		}

		msg := &dota.CNETMsg_SplitScreenUser{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SplitScreenUser {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 4: // dota.NET_Messages_net_Tick
		if c.onCNETMsg_Tick == nil {
			return nil
		}

		msg := &dota.CNETMsg_Tick{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_Tick {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 5: // dota.NET_Messages_net_StringCmd
		if c.onCNETMsg_StringCmd == nil {
			return nil
		}

		msg := &dota.CNETMsg_StringCmd{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_StringCmd {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 6: // dota.NET_Messages_net_SetConVar
		if c.onCNETMsg_SetConVar == nil {
			return nil
		}

		msg := &dota.CNETMsg_SetConVar{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SetConVar {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 7: // dota.NET_Messages_net_SignonState
		if c.onCNETMsg_SignonState == nil {
			return nil
		}

		msg := &dota.CNETMsg_SignonState{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SignonState {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 8: // dota.NET_Messages_net_SpawnGroup_Load
		if c.onCNETMsg_SpawnGroup_Load == nil {
			return nil
		}

		msg := &dota.CNETMsg_SpawnGroup_Load{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SpawnGroup_Load {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 9: // dota.NET_Messages_net_SpawnGroup_ManifestUpdate
		if c.onCNETMsg_SpawnGroup_ManifestUpdate == nil {
			return nil
		}

		msg := &dota.CNETMsg_SpawnGroup_ManifestUpdate{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SpawnGroup_ManifestUpdate {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 11: // dota.NET_Messages_net_SpawnGroup_SetCreationTick
		if c.onCNETMsg_SpawnGroup_SetCreationTick == nil {
			return nil
		}

		msg := &dota.CNETMsg_SpawnGroup_SetCreationTick{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SpawnGroup_SetCreationTick {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 12: // dota.NET_Messages_net_SpawnGroup_Unload
		if c.onCNETMsg_SpawnGroup_Unload == nil {
			return nil
		}

		msg := &dota.CNETMsg_SpawnGroup_Unload{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SpawnGroup_Unload {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 13: // dota.NET_Messages_net_SpawnGroup_LoadCompleted
		if c.onCNETMsg_SpawnGroup_LoadCompleted == nil {
			return nil
		}

		msg := &dota.CNETMsg_SpawnGroup_LoadCompleted{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_SpawnGroup_LoadCompleted {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 15: // dota.NET_Messages_net_DebugOverlay
		if c.onCNETMsg_DebugOverlay == nil {
			return nil
		}

		msg := &dota.CNETMsg_DebugOverlay{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCNETMsg_DebugOverlay {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 40: // dota.SVC_Messages_svc_ServerInfo
		if c.onCSVCMsg_ServerInfo == nil {
			return nil
		}

		msg := &dota.CSVCMsg_ServerInfo{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_ServerInfo {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 41: // dota.SVC_Messages_svc_FlattenedSerializer
		if c.onCSVCMsg_FlattenedSerializer == nil {
			return nil
		}

		msg := &dota.CSVCMsg_FlattenedSerializer{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_FlattenedSerializer {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 42: // dota.SVC_Messages_svc_ClassInfo
		if c.onCSVCMsg_ClassInfo == nil {
			return nil
		}

		msg := &dota.CSVCMsg_ClassInfo{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_ClassInfo {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 43: // dota.SVC_Messages_svc_SetPause
		if c.onCSVCMsg_SetPause == nil {
			return nil
		}

		msg := &dota.CSVCMsg_SetPause{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_SetPause {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 44: // dota.SVC_Messages_svc_CreateStringTable
		if c.onCSVCMsg_CreateStringTable == nil {
			return nil
		}

		msg := &dota.CSVCMsg_CreateStringTable{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_CreateStringTable {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 45: // dota.SVC_Messages_svc_UpdateStringTable
		if c.onCSVCMsg_UpdateStringTable == nil {
			return nil
		}

		msg := &dota.CSVCMsg_UpdateStringTable{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_UpdateStringTable {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 46: // dota.SVC_Messages_svc_VoiceInit
		if c.onCSVCMsg_VoiceInit == nil {
			return nil
		}

		msg := &dota.CSVCMsg_VoiceInit{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_VoiceInit {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 47: // dota.SVC_Messages_svc_VoiceData
		if c.onCSVCMsg_VoiceData == nil {
			return nil
		}

		msg := &dota.CSVCMsg_VoiceData{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_VoiceData {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 48: // dota.SVC_Messages_svc_Print
		if c.onCSVCMsg_Print == nil {
			return nil
		}

		msg := &dota.CSVCMsg_Print{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_Print {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 49: // dota.SVC_Messages_svc_Sounds
		if c.onCSVCMsg_Sounds == nil {
			return nil
		}

		msg := &dota.CSVCMsg_Sounds{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_Sounds {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 50: // dota.SVC_Messages_svc_SetView
		if c.onCSVCMsg_SetView == nil {
			return nil
		}

		msg := &dota.CSVCMsg_SetView{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_SetView {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 51: // dota.SVC_Messages_svc_ClearAllStringTables
		if c.onCSVCMsg_ClearAllStringTables == nil {
			return nil
		}

		msg := &dota.CSVCMsg_ClearAllStringTables{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_ClearAllStringTables {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 52: // dota.SVC_Messages_svc_CmdKeyValues
		if c.onCSVCMsg_CmdKeyValues == nil {
			return nil
		}

		msg := &dota.CSVCMsg_CmdKeyValues{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_CmdKeyValues {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 53: // dota.SVC_Messages_svc_BSPDecal
		if c.onCSVCMsg_BSPDecal == nil {
			return nil
		}

		msg := &dota.CSVCMsg_BSPDecal{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_BSPDecal {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 54: // dota.SVC_Messages_svc_SplitScreen
		if c.onCSVCMsg_SplitScreen == nil {
			return nil
		}

		msg := &dota.CSVCMsg_SplitScreen{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_SplitScreen {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 55: // dota.SVC_Messages_svc_PacketEntities
		if c.onCSVCMsg_PacketEntities == nil {
			return nil
		}

		msg := &dota.CSVCMsg_PacketEntities{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_PacketEntities {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 56: // dota.SVC_Messages_svc_Prefetch
		if c.onCSVCMsg_Prefetch == nil {
			return nil
		}

		msg := &dota.CSVCMsg_Prefetch{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_Prefetch {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 57: // dota.SVC_Messages_svc_Menu
		if c.onCSVCMsg_Menu == nil {
			return nil
		}

		msg := &dota.CSVCMsg_Menu{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_Menu {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 58: // dota.SVC_Messages_svc_GetCvarValue
		if c.onCSVCMsg_GetCvarValue == nil {
			return nil
		}

		msg := &dota.CSVCMsg_GetCvarValue{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_GetCvarValue {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 59: // dota.SVC_Messages_svc_StopSound
		if c.onCSVCMsg_StopSound == nil {
			return nil
		}

		msg := &dota.CSVCMsg_StopSound{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_StopSound {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 60: // dota.SVC_Messages_svc_PeerList
		if c.onCSVCMsg_PeerList == nil {
			return nil
		}

		msg := &dota.CSVCMsg_PeerList{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_PeerList {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 61: // dota.SVC_Messages_svc_PacketReliable
		if c.onCSVCMsg_PacketReliable == nil {
			return nil
		}

		msg := &dota.CSVCMsg_PacketReliable{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_PacketReliable {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 62: // dota.SVC_Messages_svc_HLTVStatus
		if c.onCSVCMsg_HLTVStatus == nil {
			return nil
		}

		msg := &dota.CSVCMsg_HLTVStatus{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_HLTVStatus {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 63: // dota.SVC_Messages_svc_ServerSteamID
		if c.onCSVCMsg_ServerSteamID == nil {
			return nil
		}

		msg := &dota.CSVCMsg_ServerSteamID{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_ServerSteamID {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 70: // dota.SVC_Messages_svc_FullFrameSplit
		if c.onCSVCMsg_FullFrameSplit == nil {
			return nil
		}

		msg := &dota.CSVCMsg_FullFrameSplit{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_FullFrameSplit {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 71: // dota.SVC_Messages_svc_RconServerDetails
		if c.onCSVCMsg_RconServerDetails == nil {
			return nil
		}

		msg := &dota.CSVCMsg_RconServerDetails{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_RconServerDetails {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 72: // dota.SVC_Messages_svc_UserMessage
		if c.onCSVCMsg_UserMessage == nil {
			return nil
		}

		msg := &dota.CSVCMsg_UserMessage{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_UserMessage {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 74: // dota.SVC_Messages_svc_Broadcast_Command
		if c.onCSVCMsg_Broadcast_Command == nil {
			return nil
		}

		msg := &dota.CSVCMsg_Broadcast_Command{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_Broadcast_Command {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 75: // dota.SVC_Messages_svc_HltvFixupOperatorStatus
		if c.onCSVCMsg_HltvFixupOperatorStatus == nil {
			return nil
		}

		msg := &dota.CSVCMsg_HltvFixupOperatorStatus{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCSVCMsg_HltvFixupOperatorStatus {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 101: // dota.EBaseUserMessages_UM_AchievementEvent
		if c.onCUserMessageAchievementEvent == nil {
			return nil
		}

		msg := &dota.CUserMessageAchievementEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageAchievementEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 102: // dota.EBaseUserMessages_UM_CloseCaption
		if c.onCUserMessageCloseCaption == nil {
			return nil
		}

		msg := &dota.CUserMessageCloseCaption{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageCloseCaption {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 103: // dota.EBaseUserMessages_UM_CloseCaptionDirect
		if c.onCUserMessageCloseCaptionDirect == nil {
			return nil
		}

		msg := &dota.CUserMessageCloseCaptionDirect{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageCloseCaptionDirect {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 104: // dota.EBaseUserMessages_UM_CurrentTimescale
		if c.onCUserMessageCurrentTimescale == nil {
			return nil
		}

		msg := &dota.CUserMessageCurrentTimescale{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageCurrentTimescale {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 105: // dota.EBaseUserMessages_UM_DesiredTimescale
		if c.onCUserMessageDesiredTimescale == nil {
			return nil
		}

		msg := &dota.CUserMessageDesiredTimescale{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageDesiredTimescale {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 106: // dota.EBaseUserMessages_UM_Fade
		if c.onCUserMessageFade == nil {
			return nil
		}

		msg := &dota.CUserMessageFade{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageFade {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 107: // dota.EBaseUserMessages_UM_GameTitle
		if c.onCUserMessageGameTitle == nil {
			return nil
		}

		msg := &dota.CUserMessageGameTitle{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageGameTitle {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 110: // dota.EBaseUserMessages_UM_HudMsg
		if c.onCUserMessageHudMsg == nil {
			return nil
		}

		msg := &dota.CUserMessageHudMsg{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageHudMsg {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 111: // dota.EBaseUserMessages_UM_HudText
		if c.onCUserMessageHudText == nil {
			return nil
		}

		msg := &dota.CUserMessageHudText{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageHudText {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 113: // dota.EBaseUserMessages_UM_ColoredText
		if c.onCUserMessageColoredText == nil {
			return nil
		}

		msg := &dota.CUserMessageColoredText{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageColoredText {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 114: // dota.EBaseUserMessages_UM_RequestState
		if c.onCUserMessageRequestState == nil {
			return nil
		}

		msg := &dota.CUserMessageRequestState{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageRequestState {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 115: // dota.EBaseUserMessages_UM_ResetHUD
		if c.onCUserMessageResetHUD == nil {
			return nil
		}

		msg := &dota.CUserMessageResetHUD{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageResetHUD {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 116: // dota.EBaseUserMessages_UM_Rumble
		if c.onCUserMessageRumble == nil {
			return nil
		}

		msg := &dota.CUserMessageRumble{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageRumble {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 117: // dota.EBaseUserMessages_UM_SayText
		if c.onCUserMessageSayText == nil {
			return nil
		}

		msg := &dota.CUserMessageSayText{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageSayText {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 118: // dota.EBaseUserMessages_UM_SayText2
		if c.onCUserMessageSayText2 == nil {
			return nil
		}

		msg := &dota.CUserMessageSayText2{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageSayText2 {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 119: // dota.EBaseUserMessages_UM_SayTextChannel
		if c.onCUserMessageSayTextChannel == nil {
			return nil
		}

		msg := &dota.CUserMessageSayTextChannel{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageSayTextChannel {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 120: // dota.EBaseUserMessages_UM_Shake
		if c.onCUserMessageShake == nil {
			return nil
		}

		msg := &dota.CUserMessageShake{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageShake {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 121: // dota.EBaseUserMessages_UM_ShakeDir
		if c.onCUserMessageShakeDir == nil {
			return nil
		}

		msg := &dota.CUserMessageShakeDir{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageShakeDir {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 122: // dota.EBaseUserMessages_UM_WaterShake
		if c.onCUserMessageWaterShake == nil {
			return nil
		}

		msg := &dota.CUserMessageWaterShake{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageWaterShake {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 124: // dota.EBaseUserMessages_UM_TextMsg
		if c.onCUserMessageTextMsg == nil {
			return nil
		}

		msg := &dota.CUserMessageTextMsg{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageTextMsg {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 125: // dota.EBaseUserMessages_UM_ScreenTilt
		if c.onCUserMessageScreenTilt == nil {
			return nil
		}

		msg := &dota.CUserMessageScreenTilt{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageScreenTilt {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 128: // dota.EBaseUserMessages_UM_VoiceMask
		if c.onCUserMessageVoiceMask == nil {
			return nil
		}

		msg := &dota.CUserMessageVoiceMask{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageVoiceMask {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 130: // dota.EBaseUserMessages_UM_SendAudio
		if c.onCUserMessageSendAudio == nil {
			return nil
		}

		msg := &dota.CUserMessageSendAudio{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageSendAudio {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 131: // dota.EBaseUserMessages_UM_ItemPickup
		if c.onCUserMessageItemPickup == nil {
			return nil
		}

		msg := &dota.CUserMessageItemPickup{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageItemPickup {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 132: // dota.EBaseUserMessages_UM_AmmoDenied
		if c.onCUserMessageAmmoDenied == nil {
			return nil
		}

		msg := &dota.CUserMessageAmmoDenied{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageAmmoDenied {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 134: // dota.EBaseUserMessages_UM_ShowMenu
		if c.onCUserMessageShowMenu == nil {
			return nil
		}

		msg := &dota.CUserMessageShowMenu{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageShowMenu {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 135: // dota.EBaseUserMessages_UM_CreditsMsg
		if c.onCUserMessageCreditsMsg == nil {
			return nil
		}

		msg := &dota.CUserMessageCreditsMsg{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageCreditsMsg {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 136: // dota.EBaseEntityMessages_EM_PlayJingle
		if c.onCEntityMessagePlayJingle == nil {
			return nil
		}

		msg := &dota.CEntityMessagePlayJingle{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCEntityMessagePlayJingle {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 137: // dota.EBaseEntityMessages_EM_ScreenOverlay
		if c.onCEntityMessageScreenOverlay == nil {
			return nil
		}

		msg := &dota.CEntityMessageScreenOverlay{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCEntityMessageScreenOverlay {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 138: // dota.EBaseEntityMessages_EM_RemoveAllDecals
		if c.onCEntityMessageRemoveAllDecals == nil {
			return nil
		}

		msg := &dota.CEntityMessageRemoveAllDecals{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCEntityMessageRemoveAllDecals {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 139: // dota.EBaseEntityMessages_EM_PropagateForce
		if c.onCEntityMessagePropagateForce == nil {
			return nil
		}

		msg := &dota.CEntityMessagePropagateForce{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCEntityMessagePropagateForce {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 140: // dota.EBaseEntityMessages_EM_DoSpark
		if c.onCEntityMessageDoSpark == nil {
			return nil
		}

		msg := &dota.CEntityMessageDoSpark{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCEntityMessageDoSpark {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 141: // dota.EBaseEntityMessages_EM_FixAngle
		if c.onCEntityMessageFixAngle == nil {
			return nil
		}

		msg := &dota.CEntityMessageFixAngle{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCEntityMessageFixAngle {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 142: // dota.EBaseUserMessages_UM_CloseCaptionPlaceholder
		if c.onCUserMessageCloseCaptionPlaceholder == nil {
			return nil
		}

		msg := &dota.CUserMessageCloseCaptionPlaceholder{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageCloseCaptionPlaceholder {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 143: // dota.EBaseUserMessages_UM_CameraTransition
		if c.onCUserMessageCameraTransition == nil {
			return nil
		}

		msg := &dota.CUserMessageCameraTransition{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageCameraTransition {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 144: // dota.EBaseUserMessages_UM_AudioParameter
		if c.onCUserMessageAudioParameter == nil {
			return nil
		}

		msg := &dota.CUserMessageAudioParameter{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageAudioParameter {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 150: // dota.EBaseUserMessages_UM_HapticsManagerPulse
		if c.onCUserMessageHapticsManagerPulse == nil {
			return nil
		}

		msg := &dota.CUserMessageHapticsManagerPulse{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageHapticsManagerPulse {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 151: // dota.EBaseUserMessages_UM_HapticsManagerEffect
		if c.onCUserMessageHapticsManagerEffect == nil {
			return nil
		}

		msg := &dota.CUserMessageHapticsManagerEffect{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageHapticsManagerEffect {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 153: // dota.EBaseUserMessages_UM_UpdateCssClasses
		if c.onCUserMessageUpdateCssClasses == nil {
			return nil
		}

		msg := &dota.CUserMessageUpdateCssClasses{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageUpdateCssClasses {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 154: // dota.EBaseUserMessages_UM_ServerFrameTime
		if c.onCUserMessageServerFrameTime == nil {
			return nil
		}

		msg := &dota.CUserMessageServerFrameTime{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageServerFrameTime {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 155: // dota.EBaseUserMessages_UM_LagCompensationError
		if c.onCUserMessageLagCompensationError == nil {
			return nil
		}

		msg := &dota.CUserMessageLagCompensationError{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageLagCompensationError {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 156: // dota.EBaseUserMessages_UM_RequestDllStatus
		if c.onCUserMessageRequestDllStatus == nil {
			return nil
		}

		msg := &dota.CUserMessageRequestDllStatus{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageRequestDllStatus {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 157: // dota.EBaseUserMessages_UM_RequestUtilAction
		if c.onCUserMessageRequestUtilAction == nil {
			return nil
		}

		msg := &dota.CUserMessageRequestUtilAction{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageRequestUtilAction {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 160: // dota.EBaseUserMessages_UM_RequestInventory
		if c.onCUserMessageRequestInventory == nil {
			return nil
		}

		msg := &dota.CUserMessageRequestInventory{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageRequestInventory {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 162: // dota.EBaseUserMessages_UM_RequestDiagnostic
		if c.onCUserMessageRequestDiagnostic == nil {
			return nil
		}

		msg := &dota.CUserMessageRequestDiagnostic{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCUserMessageRequestDiagnostic {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 200: // dota.EBaseGameEvents_GE_VDebugGameSessionIDEvent
		if c.onCMsgVDebugGameSessionIDEvent == nil {
			return nil
		}

		msg := &dota.CMsgVDebugGameSessionIDEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgVDebugGameSessionIDEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 201: // dota.EBaseGameEvents_GE_PlaceDecalEvent
		if c.onCMsgPlaceDecalEvent == nil {
			return nil
		}

		msg := &dota.CMsgPlaceDecalEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgPlaceDecalEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 202: // dota.EBaseGameEvents_GE_ClearWorldDecalsEvent
		if c.onCMsgClearWorldDecalsEvent == nil {
			return nil
		}

		msg := &dota.CMsgClearWorldDecalsEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgClearWorldDecalsEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 203: // dota.EBaseGameEvents_GE_ClearEntityDecalsEvent
		if c.onCMsgClearEntityDecalsEvent == nil {
			return nil
		}

		msg := &dota.CMsgClearEntityDecalsEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgClearEntityDecalsEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 204: // dota.EBaseGameEvents_GE_ClearDecalsForSkeletonInstanceEvent
		if c.onCMsgClearDecalsForSkeletonInstanceEvent == nil {
			return nil
		}

		msg := &dota.CMsgClearDecalsForSkeletonInstanceEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgClearDecalsForSkeletonInstanceEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 205: // dota.EBaseGameEvents_GE_Source1LegacyGameEventList
		if c.onCMsgSource1LegacyGameEventList == nil {
			return nil
		}

		msg := &dota.CMsgSource1LegacyGameEventList{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSource1LegacyGameEventList {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 206: // dota.EBaseGameEvents_GE_Source1LegacyListenEvents
		if c.onCMsgSource1LegacyListenEvents == nil {
			return nil
		}

		msg := &dota.CMsgSource1LegacyListenEvents{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSource1LegacyListenEvents {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 207: // dota.EBaseGameEvents_GE_Source1LegacyGameEvent
		if c.onCMsgSource1LegacyGameEvent == nil {
			return nil
		}

		msg := &dota.CMsgSource1LegacyGameEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSource1LegacyGameEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 208: // dota.EBaseGameEvents_GE_SosStartSoundEvent
		if c.onCMsgSosStartSoundEvent == nil {
			return nil
		}

		msg := &dota.CMsgSosStartSoundEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSosStartSoundEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 209: // dota.EBaseGameEvents_GE_SosStopSoundEvent
		if c.onCMsgSosStopSoundEvent == nil {
			return nil
		}

		msg := &dota.CMsgSosStopSoundEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSosStopSoundEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 210: // dota.EBaseGameEvents_GE_SosSetSoundEventParams
		if c.onCMsgSosSetSoundEventParams == nil {
			return nil
		}

		msg := &dota.CMsgSosSetSoundEventParams{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSosSetSoundEventParams {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 211: // dota.EBaseGameEvents_GE_SosSetLibraryStackFields
		if c.onCMsgSosSetLibraryStackFields == nil {
			return nil
		}

		msg := &dota.CMsgSosSetLibraryStackFields{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSosSetLibraryStackFields {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 212: // dota.EBaseGameEvents_GE_SosStopSoundEventHash
		if c.onCMsgSosStopSoundEventHash == nil {
			return nil
		}

		msg := &dota.CMsgSosStopSoundEventHash{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgSosStopSoundEventHash {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 300: // dota.CitadelUserMessageIds_k_EUserMsg_Damage
		if c.onCCitadelUserMessage_Damage == nil {
			return nil
		}

		msg := &dota.CCitadelUserMessage_Damage{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMessage_Damage {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 303: // dota.CitadelUserMessageIds_k_EUserMsg_MapPing
		if c.onCCitadelUserMsg_MapPing == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_MapPing{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_MapPing {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 304: // dota.CitadelUserMessageIds_k_EUserMsg_TeamRewards
		if c.onCCitadelUserMsg_TeamRewards == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_TeamRewards{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_TeamRewards {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 308: // dota.CitadelUserMessageIds_k_EUserMsg_TriggerDamageFlash
		if c.onCCitadelUserMsg_TriggerDamageFlash == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_TriggerDamageFlash{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_TriggerDamageFlash {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 309: // dota.CitadelUserMessageIds_k_EUserMsg_AbilitiesChanged
		if c.onCCitadelUserMsg_AbilitiesChanged == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_AbilitiesChanged{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_AbilitiesChanged {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 310: // dota.CitadelUserMessageIds_k_EUserMsg_RecentDamageSummary
		if c.onCCitadelUserMsg_RecentDamageSummary == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_RecentDamageSummary{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_RecentDamageSummary {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 311: // dota.CitadelUserMessageIds_k_EUserMsg_SpectatorTeamChanged
		if c.onCCitadelUserMsg_SpectatorTeamChanged == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_SpectatorTeamChanged{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_SpectatorTeamChanged {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 312: // dota.CitadelUserMessageIds_k_EUserMsg_ChatWheel
		if c.onCCitadelUserMsg_ChatWheel == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_ChatWheel{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_ChatWheel {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 313: // dota.CitadelUserMessageIds_k_EUserMsg_GoldHistory
		if c.onCCitadelUserMsg_GoldHistory == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_GoldHistory{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_GoldHistory {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 314: // dota.CitadelUserMessageIds_k_EUserMsg_ChatMsg
		if c.onCCitadelUserMsg_ChatMsg == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_ChatMsg{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_ChatMsg {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 315: // dota.CitadelUserMessageIds_k_EUserMsg_QuickResponse
		if c.onCCitadelUserMsg_QuickResponse == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_QuickResponse{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_QuickResponse {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 316: // dota.CitadelUserMessageIds_k_EUserMsg_PostMatchDetails
		if c.onCCitadelUserMsg_PostMatchDetails == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_PostMatchDetails{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_PostMatchDetails {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 317: // dota.CitadelUserMessageIds_k_EUserMsg_ChatEvent
		if c.onCCitadelUserMsg_ChatEvent == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_ChatEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_ChatEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 318: // dota.CitadelUserMessageIds_k_EUserMsg_AbilityInterrupted
		if c.onCCitadelUserMsg_AbilityInterrupted == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_AbilityInterrupted{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_AbilityInterrupted {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 319: // dota.CitadelUserMessageIds_k_EUserMsg_HeroKilled
		if c.onCCitadelUserMsg_HeroKilled == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_HeroKilled{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_HeroKilled {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 320: // dota.CitadelUserMessageIds_k_EUserMsg_ReturnIdol
		if c.onCCitadelUserMsg_ReturnIdol == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_ReturnIdol{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_ReturnIdol {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 321: // dota.CitadelUserMessageIds_k_EUserMsg_SetClientCameraAngles
		if c.onCCitadelUserMsg_SetClientCameraAngles == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_SetClientCameraAngles{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_SetClientCameraAngles {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 322: // dota.CitadelUserMessageIds_k_EUserMsg_MapLine
		if c.onCCitadelUserMsg_MapLine == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_MapLine{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_MapLine {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 323: // dota.CitadelUserMessageIds_k_EUserMsg_BulletHit
		if c.onCCitadelUserMessage_BulletHit == nil {
			return nil
		}

		msg := &dota.CCitadelUserMessage_BulletHit{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMessage_BulletHit {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 324: // dota.CitadelUserMessageIds_k_EUserMsg_ObjectiveMask
		if c.onCCitadelUserMessage_ObjectiveMask == nil {
			return nil
		}

		msg := &dota.CCitadelUserMessage_ObjectiveMask{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMessage_ObjectiveMask {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 325: // dota.CitadelUserMessageIds_k_EUserMsg_ModifierApplied
		if c.onCCitadelUserMessage_ModifierApplied == nil {
			return nil
		}

		msg := &dota.CCitadelUserMessage_ModifierApplied{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMessage_ModifierApplied {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 326: // dota.CitadelUserMessageIds_k_EUserMsg_CameraController
		if c.onCCitadelUserMsg_CameraController == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_CameraController{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_CameraController {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 327: // dota.CitadelUserMessageIds_k_EUserMsg_AuraModifierApplied
		if c.onCCitadelUserMessage_AuraModifierApplied == nil {
			return nil
		}

		msg := &dota.CCitadelUserMessage_AuraModifierApplied{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMessage_AuraModifierApplied {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 329: // dota.CitadelUserMessageIds_k_EUserMsg_ObstructedShotFired
		if c.onCCitadelUserMsg_ObstructedShotFired == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_ObstructedShotFired{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_ObstructedShotFired {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 330: // dota.CitadelUserMessageIds_k_EUserMsg_AbilityLateFailure
		if c.onCCitadelUserMsg_AbilityLateFailure == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_AbilityLateFailure{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_AbilityLateFailure {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 331: // dota.CitadelUserMessageIds_k_EUserMsg_AbilityPing
		if c.onCCitadelUserMsg_AbilityPing == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_AbilityPing{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_AbilityPing {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 332: // dota.CitadelUserMessageIds_k_EUserMsg_PostProcessingAnim
		if c.onCCitadelUserMsg_PostProcessingAnim == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_PostProcessingAnim{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_PostProcessingAnim {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 333: // dota.CitadelUserMessageIds_k_EUserMsg_DeathReplayData
		if c.onCCitadelUserMsg_DeathReplayData == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_DeathReplayData{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_DeathReplayData {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 334: // dota.CitadelUserMessageIds_k_EUserMsg_PlayerLifetimeStatInfo
		if c.onCCitadelUserMsg_PlayerLifetimeStatInfo == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_PlayerLifetimeStatInfo{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_PlayerLifetimeStatInfo {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 336: // dota.CitadelUserMessageIds_k_EUserMsg_ForceShopClosed
		if c.onCCitadelUserMsg_ForceShopClosed == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_ForceShopClosed{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_ForceShopClosed {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 337: // dota.CitadelUserMessageIds_k_EUserMsg_StaminaDrained
		if c.onCCitadelUserMsg_StaminaDrained == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_StaminaDrained{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_StaminaDrained {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 338: // dota.CitadelUserMessageIds_k_EUserMsg_AbilityNotify
		if c.onCCitadelUserMessage_AbilityNotify == nil {
			return nil
		}

		msg := &dota.CCitadelUserMessage_AbilityNotify{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMessage_AbilityNotify {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 339: // dota.CitadelUserMessageIds_k_EUserMsg_GetDamageStatsResponse
		if c.onCCitadelUserMsg_GetDamageStatsResponse == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_GetDamageStatsResponse{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_GetDamageStatsResponse {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 340: // dota.CitadelUserMessageIds_k_EUserMsg_ParticipantStartSoundEvent
		if c.onCCitadelUserMsg_ParticipantStartSoundEvent == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_ParticipantStartSoundEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_ParticipantStartSoundEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 341: // dota.CitadelUserMessageIds_k_EUserMsg_ParticipantStopSoundEvent
		if c.onCCitadelUserMsg_ParticipantStopSoundEvent == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_ParticipantStopSoundEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_ParticipantStopSoundEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 342: // dota.CitadelUserMessageIds_k_EUserMsg_ParticipantStopSoundEventHash
		if c.onCCitadelUserMsg_ParticipantStopSoundEventHash == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_ParticipantStopSoundEventHash{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_ParticipantStopSoundEventHash {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 343: // dota.CitadelUserMessageIds_k_EUserMsg_ParticipantSetSoundEventParams
		if c.onCCitadelUserMsg_ParticipantSetSoundEventParams == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_ParticipantSetSoundEventParams{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_ParticipantSetSoundEventParams {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 344: // dota.CitadelUserMessageIds_k_EUserMsg_ParticipantSetLibraryStackFields
		if c.onCCitadelUserMsg_ParticipantSetLibraryStackFields == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_ParticipantSetLibraryStackFields{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_ParticipantSetLibraryStackFields {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 347: // dota.CitadelUserMessageIds_k_EUserMsg_BossKilled
		if c.onCCitadelUserMsg_BossKilled == nil {
			return nil
		}

		msg := &dota.CCitadelUserMsg_BossKilled{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCCitadelUserMsg_BossKilled {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 450: // dota.ECitadelGameEvents_GE_FireBullets
		if c.onCMsgFireBullets == nil {
			return nil
		}

		msg := &dota.CMsgFireBullets{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgFireBullets {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 451: // dota.ECitadelGameEvents_GE_PlayerAnimEvent
		if c.onCMsgPlayerAnimEvent == nil {
			return nil
		}

		msg := &dota.CMsgPlayerAnimEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgPlayerAnimEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 458: // dota.ECitadelGameEvents_GE_ParticleSystemManager
		if c.onCMsgParticleSystemManager == nil {
			return nil
		}

		msg := &dota.CMsgParticleSystemManager{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgParticleSystemManager {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 459: // dota.ECitadelGameEvents_GE_ScreenTextPretty
		if c.onCMsgScreenTextPretty == nil {
			return nil
		}

		msg := &dota.CMsgScreenTextPretty{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgScreenTextPretty {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 460: // dota.ECitadelGameEvents_GE_ServerRequestedTracer
		if c.onCMsgServerRequestedTracer == nil {
			return nil
		}

		msg := &dota.CMsgServerRequestedTracer{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgServerRequestedTracer {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 461: // dota.ECitadelGameEvents_GE_BulletImpact
		if c.onCMsgBulletImpact == nil {
			return nil
		}

		msg := &dota.CMsgBulletImpact{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgBulletImpact {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 462: // dota.ECitadelGameEvents_GE_EnableSatVolumesEvent
		if c.onCMsgEnableSatVolumesEvent == nil {
			return nil
		}

		msg := &dota.CMsgEnableSatVolumesEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgEnableSatVolumesEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 463: // dota.ECitadelGameEvents_GE_PlaceSatVolumeEvent
		if c.onCMsgPlaceSatVolumeEvent == nil {
			return nil
		}

		msg := &dota.CMsgPlaceSatVolumeEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgPlaceSatVolumeEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 464: // dota.ECitadelGameEvents_GE_DisableSatVolumesEvent
		if c.onCMsgDisableSatVolumesEvent == nil {
			return nil
		}

		msg := &dota.CMsgDisableSatVolumesEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgDisableSatVolumesEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	case 465: // dota.ECitadelGameEvents_GE_RemoveSatVolumeEvent
		if c.onCMsgRemoveSatVolumeEvent == nil {
			return nil
		}

		msg := &dota.CMsgRemoveSatVolumeEvent{}
		c.pb.SetBuf(buf)
		if err := c.pb.Unmarshal(msg); err != nil {
			return err
		}

		for _, fn := range c.onCMsgRemoveSatVolumeEvent {
			if err := fn(msg); err != nil {
				return err
			}
		}

		return nil

	}

	return nil
}
