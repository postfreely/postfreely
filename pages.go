/*
 * Copyright © 2018-2019, 2021 Musing Studio LLC.
 *
 * This file is part of WriteFreely.
 *
 * WriteFreely is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License, included
 * in the LICENSE file in this source code package.
 */

package postfreely

import (
	"database/sql"
	"github.com/postfreely/postfreely/config"
	"time"
)

var defaultPageUpdatedTime = time.Date(2018, 11, 8, 12, 0, 0, 0, time.Local)

func getAboutPage(app *App) (*instanceContent, error) {
	c, err := app.db.GetDynamicContent("about")
	if err != nil {
		return nil, err
	}
	if c == nil {
		c = &instanceContent{
			ID:      "about",
			Type:    "page",
			Content: defaultAboutPage(app.Config()),
		}
	}
	if !c.Title.Valid {
		c.Title = defaultAboutTitle(app.Config())
	}
	return c, nil
}

func defaultAboutTitle(cfg *config.Config) sql.NullString {
	return sql.NullString{String: "About " + cfg.App.SiteName, Valid: true}
}

func getContactPage(app *App) (*instanceContent, error) {
	c, err := app.db.GetDynamicContent("contact")
	if err != nil {
		return nil, err
	}
	if c == nil {
		c = &instanceContent{
			ID:      "contact",
			Type:    "page",
			Content: defaultContactPage(app),
		}
	}
	if !c.Title.Valid {
		c.Title = defaultContactTitle()
	}
	return c, nil
}

func defaultContactTitle() sql.NullString {
	return sql.NullString{String: "Contact Us", Valid: true}
}

func getPrivacyPage(app *App) (*instanceContent, error) {
	c, err := app.db.GetDynamicContent("privacy")
	if err != nil {
		return nil, err
	}
	if c == nil {
		c = &instanceContent{
			ID:      "privacy",
			Type:    "page",
			Content: defaultPrivacyPolicy(app.Config()),
			Updated: defaultPageUpdatedTime,
		}
	}
	if !c.Title.Valid {
		c.Title = defaultPrivacyTitle()
	}
	return c, nil
}

func defaultPrivacyTitle() sql.NullString {
	return sql.NullString{String: "Privacy Policy", Valid: true}
}

func defaultAboutPage(cfg *config.Config) string {
	if cfg.App.Federation {
		return `_` + cfg.App.SiteName + `_ is an interconnected place for you to write and publish, powered by [` + serverSoftware + `](` + softwareURL + `) and ActivityPub.`
	}
	return `_` + cfg.App.SiteName + `_ is a place for you to write and publish, powered by [` + serverSoftware + `](` + softwareURL + `).`
}

func defaultContactPage(app *App) string {
	c, err := app.db.GetCollectionByID(1)
	if err != nil {
		return ""
	}
	return `_` + app.Config().App.SiteName + `_ is administered by: [**` + c.Alias + `**](/` + c.Alias + `/).

Contact them at this email address: _EMAIL GOES HERE_.

You can also reach them here...`
}

func defaultPrivacyPolicy(cfg *config.Config) string {
	return `[` + serverSoftware + `](` + softwareURL + `), the software that powers this site, is built to enforce your right to privacy by default.

It retains as little data about you as possible, not even requiring an email address to sign up. However, if you _do_ give us your email address, it is stored encrypted in our database. We salt and hash your account's password.

We store log files, or data about what happens on our servers. We also use cookies to keep you logged into your account.

Beyond this, it's important that you trust whoever runs **` + cfg.App.SiteName + `**. Software can only do so much to protect you -- your level of privacy protections will ultimately fall on the humans that run this particular service.`
}

func getLandingBanner(app *App) (*instanceContent, error) {
	c, err := app.db.GetDynamicContent("landing-banner")
	if err != nil {
		return nil, err
	}
	if c == nil {
		c = &instanceContent{
			ID:      "landing-banner",
			Type:    "section",
			Content: defaultLandingBanner(app.Config()),
			Updated: defaultPageUpdatedTime,
		}
	}
	return c, nil
}

func getLandingBody(app *App) (*instanceContent, error) {
	c, err := app.db.GetDynamicContent("landing-body")
	if err != nil {
		return nil, err
	}
	if c == nil {
		c = &instanceContent{
			ID:      "landing-body",
			Type:    "section",
			Content: defaultLandingBody(app.Config()),
			Updated: defaultPageUpdatedTime,
		}
	}
	return c, nil
}

func defaultLandingBanner(cfg *config.Config) string {
	if cfg.App.Federation {
		return "# Start your blog in the fediverse"
	}
	return "# Start your blog"
}

func defaultLandingBody(cfg *config.Config) string {
	if cfg.App.Federation {
		return `## Join the Fediverse

The fediverse is a large network of platforms that all speak a common language. Imagine if you could reply to Instagram posts from Twitter, or interact with your favorite Medium blogs from Facebook -- federated alternatives like [PixelFed](https://pixelfed.org), [Mastodon](https://joinmastodon.org), and ` + serverSoftware + ` enable you to do these types of things.

<div style="text-align:center">
	<iframe style="width: 560px; height: 315px; max-width: 100%;" sandbox="allow-same-origin allow-scripts" src="https://video.writeas.org/videos/embed/cc55e615-d204-417c-9575-7b57674cc6f3" frameborder="0" allowfullscreen></iframe>
</div>

## Write More Socially

` + serverSoftware + ` can communicate with other federated platforms like Mastodon, so people can follow your blogs, bookmark their favorite posts, and boost them to their followers. Sign up above to create a blog and join the fediverse.`
	}
	return ""
}

func getReaderSection(app *App) (*instanceContent, error) {
	c, err := app.db.GetDynamicContent("reader")
	if err != nil {
		return nil, err
	}
	if c == nil {
		c = &instanceContent{
			ID:      "reader",
			Type:    "section",
			Content: defaultReaderBanner(app.Config()),
			Updated: defaultPageUpdatedTime,
		}
	}
	if !c.Title.Valid {
		c.Title = defaultReaderTitle(app.Config())
	}
	return c, nil
}

func defaultReaderTitle(cfg *config.Config) sql.NullString {
	return sql.NullString{String: "Reader", Valid: true}
}

func defaultReaderBanner(cfg *config.Config) string {
	return "Read the latest posts from " + cfg.App.SiteName + "."
}
