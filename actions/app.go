package actions

import (
	"github.com/JewlyTwin/be_booking_sign/actions/handlers"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	forcessl "github.com/gobuffalo/mw-forcessl"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/unrolled/secure"

	// "github.com/JewlyTwin/practice/models"
	"github.com/gobuffalo/buffalo-pop/pop/popmw"
	contenttype "github.com/gobuffalo/mw-contenttype"
	"github.com/gobuffalo/x/sessions"
	"github.com/rs/cors"

	"github.com/JewlyTwin/be_booking_sign/models"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			SessionStore: sessions.Null{},
			PreWares: []buffalo.PreWare{
				cors.AllowAll().Handler,
			},
			SessionName: "_practice_session",
		})

		// Automatically redirect to SSL
		app.Use(forceSSL())

		// Log request parameters (filters apply).
		app.Use(paramlogger.ParameterLogger)

		// Set the request content type to JSON
		app.Use(contenttype.Set("application/json"))

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.Connection)
		// Remove to disable this.
		app.Use(popmw.Transaction(models.DB))

		//Authentication
		app.POST("/login", handlers.Login)
		app.POST("/register", handlers.Register)
		//sign
		app.GET("/allsign", handlers.GetAllSign)
		app.POST("/addsign", handlers.AddSign)
		app.POST("/addbooking", handlers.AddBooking)
		app.GET("/booking/{page}/{order}", handlers.GetPaginateUser)
		app.POST("/deletesign", handlers.DeleteSign)
		app.POST("/updatesign", handlers.UpdateSign)
		app.GET("/sign/{id}", handlers.GetSignById)
		//booking
		app.POST("/addbooking", handlers.AddBooking)
		// app.POST("/deletebooking", handlers.DeleteBooking)
		app.GET("/getbookingdays/{id}", handlers.GetBookingDayBySign)
		app.POST("/login", handlers.Login)
		//user
		app.GET("/user", handlers.GetUserById)
		//admin
		app.GET("/admin/booking/{page}", handlers.GetPaginateAdmin)
		app.POST("/admin/booking/approve", handlers.ApproveBooking)
		app.POST("/admin/booking/reject", handlers.RejectBooking)
		//test
		app.GET("/sendmail", handlers.SendMail)

	}

	return app
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}
