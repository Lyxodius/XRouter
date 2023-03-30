package xrouter

import "github.com/gorilla/mux"

func NewRouter() *XRouter {
	return &XRouter{namedRoutes: make(map[string]*XRoute)}
}

type XRouter struct {
	mux.Router
	namedRoutes map[string]*XRoute
}

type XRoute struct {
	mux.Route

	requiredPermission Permission
}

func (r *XRoute) Permission(permission Permission) {
	r.requiredPermission = permission
}

type Permission struct {
	Area   permissionArea
	Action permissionAction
}

type permissions struct {
	ViewTests                     Permission
	StartTests                    Permission
	StopAnyTests                  Permission
	ModifyOwnTests                Permission
	ModifyAnyTests                Permission
	ViewOutboundConnections       Permission
	ModifyOutboundConnections     Permission
	DeleteAnyOutboundConnections  Permission
	ViewInboundConnections        Permission
	ModifyInboundConnections      Permission
	DeleteAnyInboundConnections   Permission
	ViewSubscriberMobile          Permission
	ModifySubscriberMobile        Permission
	DeleteAnySubscriberMobile     Permission
	ViewDictionary                Permission
	ModifyDictionary              Permission
	DeleteAnyDictionary           Permission
	ViewScreens                   Permission
	ModifyScreens                 Permission
	DeleteAnyScreens              Permission
	ViewEvaluationParameters      Permission
	ModifyEvaluationParameters    Permission
	DeleteAnyEvaluationParameters Permission
	ViewObjects                   Permission
	ModifyObjects                 Permission
	DeleteAnyObjects              Permission
}

var PermissionType = permissions{
	ViewTests: Permission{
		Area:   PermissionArea.Tests,
		Action: PermissionAction.View,
	},
	StartTests: Permission{
		Area:   PermissionArea.Tests,
		Action: PermissionAction.Start,
	},
	StopAnyTests: Permission{
		Area:   PermissionArea.Tests,
		Action: PermissionAction.StopAny,
	},
	ModifyOwnTests: Permission{
		Area:   PermissionArea.Tests,
		Action: PermissionAction.ModifyOwn,
	},
	ModifyAnyTests: Permission{
		Area:   PermissionArea.Tests,
		Action: PermissionAction.ModifyAny,
	},
	ViewOutboundConnections: Permission{
		Area:   PermissionArea.OutboundConnections,
		Action: PermissionAction.View,
	},
	ModifyOutboundConnections: Permission{
		Area:   PermissionArea.OutboundConnections,
		Action: PermissionAction.Modify,
	},
	DeleteAnyOutboundConnections: Permission{
		Area:   PermissionArea.OutboundConnections,
		Action: PermissionAction.DeleteAny,
	},
	ViewInboundConnections: Permission{
		Area:   PermissionArea.InboundConnections,
		Action: PermissionAction.View,
	},
	ModifyInboundConnections: Permission{
		Area:   PermissionArea.InboundConnections,
		Action: PermissionAction.Modify,
	},
	DeleteAnyInboundConnections: Permission{
		Area:   PermissionArea.InboundConnections,
		Action: PermissionAction.DeleteAny,
	},
	ViewSubscriberMobile: Permission{
		Area:   PermissionArea.SubscriberMobile,
		Action: PermissionAction.View,
	},
	ModifySubscriberMobile: Permission{
		Area:   PermissionArea.SubscriberMobile,
		Action: PermissionAction.Modify,
	},
	DeleteAnySubscriberMobile: Permission{
		Area:   PermissionArea.SubscriberMobile,
		Action: PermissionAction.DeleteAny,
	},
	ViewDictionary: Permission{
		Area:   PermissionArea.Dictionary,
		Action: PermissionAction.View,
	},
	ModifyDictionary: Permission{
		Area:   PermissionArea.Dictionary,
		Action: PermissionAction.Modify,
	},
	DeleteAnyDictionary: Permission{
		Area:   PermissionArea.Dictionary,
		Action: PermissionAction.DeleteAny,
	},
	ViewScreens: Permission{
		Area:   PermissionArea.Screens,
		Action: PermissionAction.View,
	},
	ModifyScreens: Permission{
		Area:   PermissionArea.Screens,
		Action: PermissionAction.Modify,
	},
	DeleteAnyScreens: Permission{
		Area:   PermissionArea.Screens,
		Action: PermissionAction.DeleteAny,
	},
	ViewEvaluationParameters: Permission{
		Area:   PermissionArea.EvaluationParameters,
		Action: PermissionAction.View,
	},
	ModifyEvaluationParameters: Permission{
		Area:   PermissionArea.EvaluationParameters,
		Action: PermissionAction.Modify,
	},
	DeleteAnyEvaluationParameters: Permission{
		Area:   PermissionArea.EvaluationParameters,
		Action: PermissionAction.DeleteAny,
	},
	ViewObjects: Permission{
		Area:   PermissionArea.Objects,
		Action: PermissionAction.View,
	},
	ModifyObjects: Permission{
		Area:   PermissionArea.Objects,
		Action: PermissionAction.Modify,
	},
	DeleteAnyObjects: Permission{
		Area:   PermissionArea.Objects,
		Action: PermissionAction.DeleteAny,
	},
}

var allPermissions = []Permission{
	PermissionType.ViewTests,
	PermissionType.StartTests,
	PermissionType.StopAnyTests,
	PermissionType.ModifyOwnTests,
	PermissionType.ModifyAnyTests,
	PermissionType.ViewOutboundConnections,
	PermissionType.ModifyOutboundConnections,
	PermissionType.DeleteAnyOutboundConnections,
	PermissionType.ViewInboundConnections,
	PermissionType.ModifyInboundConnections,
	PermissionType.DeleteAnyInboundConnections,
	PermissionType.ViewSubscriberMobile,
	PermissionType.ModifySubscriberMobile,
	PermissionType.DeleteAnySubscriberMobile,
	PermissionType.ViewDictionary,
	PermissionType.ModifyDictionary,
	PermissionType.DeleteAnyDictionary,
	PermissionType.ViewScreens,
	PermissionType.ModifyScreens,
	PermissionType.DeleteAnyScreens,
	PermissionType.ViewEvaluationParameters,
	PermissionType.ModifyEvaluationParameters,
	PermissionType.DeleteAnyEvaluationParameters,
	PermissionType.ViewObjects,
	PermissionType.ModifyObjects,
	PermissionType.DeleteAnyObjects,
}

type permissionArea struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type permissionAreas struct {
	Tests                permissionArea
	OutboundConnections  permissionArea
	InboundConnections   permissionArea
	SubscriberMobile     permissionArea
	Dictionary           permissionArea
	Screens              permissionArea
	EvaluationParameters permissionArea
	Objects              permissionArea
}

type permissionAction struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type permissionActions struct {
	Create               permissionAction
	EditOwn              permissionAction
	EditAny              permissionAction
	Start                permissionAction
	StopOwn              permissionAction
	StopAny              permissionAction
	DeleteOwn            permissionAction
	DeleteAny            permissionAction
	View                 permissionAction
	ViewInternalComments permissionAction
	Modify               permissionAction
	ModifyAny            permissionAction
	ModifyOwn            permissionAction
}

var PermissionArea = permissionAreas{
	Tests: permissionArea{
		Name:  "Tests",
		Value: "1",
	},
	OutboundConnections: permissionArea{
		Name:  "OutboundConnections",
		Value: "2",
	},
	InboundConnections: permissionArea{
		Name:  "InboundConnections",
		Value: "3",
	},
	SubscriberMobile: permissionArea{
		Name:  "SubscriberMobile",
		Value: "4",
	},
	Dictionary: permissionArea{
		Name:  "Dictionary",
		Value: "5",
	},
	Screens: permissionArea{
		Name:  "Screens",
		Value: "6",
	},
	EvaluationParameters: permissionArea{
		Name:  "EvaluationParameters",
		Value: "7",
	},
	Objects: permissionArea{
		Name:  "Objects",
		Value: "8",
	},
}

var PermissionAction = permissionActions{
	Create: permissionAction{
		Name:  "Create",
		Value: "A",
	},
	EditOwn: permissionAction{
		Name:  "EditOwn",
		Value: "B",
	},
	EditAny: permissionAction{
		Name:  "EditAny",
		Value: "C",
	},
	Start: permissionAction{
		Name:  "Start",
		Value: "D",
	},
	StopOwn: permissionAction{
		Name:  "StopOwn",
		Value: "E",
	},
	StopAny: permissionAction{
		Name:  "StopAny",
		Value: "F",
	},
	DeleteOwn: permissionAction{
		Name:  "DeleteOwn",
		Value: "H",
	},
	DeleteAny: permissionAction{
		Name:  "DeleteAny",
		Value: "I",
	},
	View: permissionAction{
		Name:  "View",
		Value: "K",
	},
	ViewInternalComments: permissionAction{
		Name:  "ViewInternalComments",
		Value: "L",
	},
	Modify: permissionAction{
		Name:  "Modify",
		Value: "M",
	},
	ModifyAny: permissionAction{
		Name:  "ModifyAny",
		Value: "N",
	},
	ModifyOwn: permissionAction{
		Name:  "ModifyOwn",
		Value: "O",
	},
}

func GetPermissionByValue(value string) *Permission {
	if len(value) != 2 {
		return nil
	}
	areaRune := string(value[0])
	actionRune := string(value[1])

	for _, perm := range allPermissions {
		if perm.Area.Value == areaRune && perm.Action.Value == actionRune {
			return &perm
		}
	}
	return nil
}
