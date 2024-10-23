// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package marketing

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/ctx"
	models "github.com/onsonr/sonr/internal/orm/marketing"
	"github.com/onsonr/sonr/pkg/nebula/components/marketing/sections"
	"github.com/onsonr/sonr/pkg/nebula/global/styles"
	"log"
)

// ╭───────────────────────────────────────────────────────────╮
// │                  Home Routes - Marketing                  │
// ╰───────────────────────────────────────────────────────────╯

func HomeRoute(c echo.Context) error {
	s, err := ctx.GetHWAYContext(c)
	if err != nil {
		return err
	}
	log.Printf("Session ID: %s", s.ID())
	return ctx.RenderTempl(c, View())
}

// ╭───────────────────────────────────────────────────────────╮
// │                  Static Content Values                    │
// ╰───────────────────────────────────────────────────────────╯

var hero = &models.Hero{
	TitleFirst:      "Simplified",
	TitleEmphasis:   "self-custody",
	TitleSecond:     "for everyone",
	Subtitle:        "Sonr is a modern re-imagination of online user identity, empowering users to take ownership of their digital footprint and unlocking a new era of self-sovereignty.",
	PrimaryButton:   &models.Button{Text: "Get Started", Href: "/register"},
	SecondaryButton: &models.Button{Text: "Learn More", Href: "/about"},
	Image: &models.Image{
		Src:    "https://cdn.sonr.id/img/hero-clipped.svg",
		Width:  "500",
		Height: "500",
	},
	Stats: []*models.Stat{
		{Value: "476", Label: "Assets packed with power beyond your imagiation.", Denom: "K"},
		{Value: "1.44", Label: "Assets packed with power beyond your imagination.", Denom: "K"},
		{Value: "1.5", Label: "Assets packed with power beyond your imagination.", Denom: "M+"},
		{Value: "750", Label: "Assets packed with power beyond your imagination.", Denom: "K"},
	},
}

var highlights = &models.Highlights{
	Heading:  "The Internet Rebuilt for You",
	Subtitle: "Sonr is a comprehensive system for Identity Management which proteects users across their digital personas while providing Developers a cost-effective solution for decentralized authentication.",
	Features: []*models.Feature{
		{
			Title: "∞ Factor Auth",
			Desc:  "Sonr is designed to work across all platforms and devices, building a encrypted and anonymous identity layer for each user on the internet.",
			Icon:  nil,
		},
		{
			Title: "Control Your Data",
			Desc:  "Sonr leverages advanced cryptography to permit facilitating Wallet Operations directly on-chain, without the need for a centralized server.",
			Icon:  nil,
		},
		{
			Title: "Crypto Enabled",
			Desc:  "Sonr follows the latest specifications from W3C, DIF, and ICF to essentially have an Interchain-Connected, Smart Account System - seamlessly authenticated with PassKeys.",
			Icon:  nil,
		},
		{
			Title: "Works Everywhere",
			Desc:  "Sonr anonymously associates your online identities with a Quantum-Resistant Vault which only you can access.",
			Icon:  nil,
		},
	},
}

var mission = &models.Mission{
	Eyebrow:  "L1 Blockchain",
	Heading:  "The Protocol for Decentralized Identity & Authentication",
	Subtitle: "We're creating the Global Standard for Decentralized Identity. Authenticate users with PassKeys, Issue Crypto Wallets, Build Payment flows, Send Encrypted Messages - all on a single platform.",
	Experience: &models.Feature{
		Title: "Less is More",
		Desc:  "Sonr is a comprehensive system for Identity Management which proteects users across their digital personas while providing Developers a cost-effective solution for decentralized authentication.",
		Icon:  nil,
	},
	Compliance: &models.Feature{
		Title: "Works where there's Internet",
		Desc:  "Sonr is designed to work across all platforms and devices, building a encrypted and anonymous identity layer for each user on the internet.",
		Icon:  nil,
	},
	Interoperability: &models.Feature{
		Title: "Made in the USA",
		Desc:  "Sonr follows the latest specifications from W3C, DIF, and ICF to essentially have an Interchain-Connected, Smart Account System - seamlessly authenticated with PassKeys.",
		Icon:  nil,
	},
}

// ╭─────────────────────────────────────────────────────────╮
// │                  Final Rendering                        │
// ╰─────────────────────────────────────────────────────────╯

// View renders the home page
func View() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			templ_7745c5c3_Err = sections.Header().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = sections.Hero(hero).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = sections.Highlights(highlights).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = sections.Mission(mission).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = sections.Architecture().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = sections.Lowlights().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = sections.CallToAction().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = sections.Footer().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = styles.LayoutNoBody("Sonr.ID", true).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
