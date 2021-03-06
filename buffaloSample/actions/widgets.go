package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/jimxshaw/testing/buffaloSample/models"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Widget)
// DB Table: Plural (widgets)
// Resource: Plural (Widgets)
// Path: Plural (/widgets)
// View Template Folder: Plural (/templates/widgets/)

// WidgetsResource is the resource for the widget model
type WidgetsResource struct {
	buffalo.Resource
}

// List gets all Widgets. This function is mapped to the path
// GET /widgets
func (v WidgetsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	widgets := &models.Widgets{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Widgets from the DB
	if err := q.All(widgets); err != nil {
		return errors.WithStack(err)
	}

	// Make Widgets available inside the html template
	c.Set("widgets", widgets)

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.HTML("widgets/index.html"))
}

// Show gets the data for one Widget. This function is mapped to
// the path GET /widgets/{widget_id}
func (v WidgetsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Widget
	widget := &models.Widget{}

	// To find the Widget the parameter widget_id is used.
	if err := tx.Find(widget, c.Param("widget_id")); err != nil {
		return c.Error(404, err)
	}

	// Make widget available inside the html template
	c.Set("widget", widget)

	return c.Render(200, r.HTML("widgets/show.html"))
}

// New renders the form for creating a new Widget.
// This function is mapped to the path GET /widgets/new
func (v WidgetsResource) New(c buffalo.Context) error {
	// Make widget available inside the html template
	c.Set("widget", &models.Widget{})

	return c.Render(200, r.HTML("widgets/new.html"))
}

// Create adds a Widget to the DB. This function is mapped to the
// path POST /widgets
func (v WidgetsResource) Create(c buffalo.Context) error {
	// Allocate an empty Widget
	widget := &models.Widget{}

	// Bind widget to the html form elements
	if err := c.Bind(widget); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(widget)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make widget available inside the html template
		c.Set("widget", widget)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("widgets/new.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Widget was created successfully")

	// and redirect to the widgets index page
	return c.Redirect(302, "/widgets/%s", widget.ID)
}

// Edit renders a edit form for a widget. This function is
// mapped to the path GET /widgets/{widget_id}/edit
func (v WidgetsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Widget
	widget := &models.Widget{}

	if err := tx.Find(widget, c.Param("widget_id")); err != nil {
		return c.Error(404, err)
	}

	// Make widget available inside the html template
	c.Set("widget", widget)
	return c.Render(200, r.HTML("widgets/edit.html"))
}

// Update changes a widget in the DB. This function is mapped to
// the path PUT /widgets/{widget_id}
func (v WidgetsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Widget
	widget := &models.Widget{}

	if err := tx.Find(widget, c.Param("widget_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Widget to the html form elements
	if err := c.Bind(widget); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(widget)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make widget available inside the html template
		c.Set("widget", widget)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("widgets/edit.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Widget was updated successfully")

	// and redirect to the widgets index page
	return c.Redirect(302, "/widgets/%s", widget.ID)
}

// Destroy deletes a widget from the DB. This function is mapped
// to the path DELETE /widgets/{widget_id}
func (v WidgetsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Widget
	widget := &models.Widget{}

	// To find the Widget the parameter widget_id is used.
	if err := tx.Find(widget, c.Param("widget_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(widget); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Widget was destroyed successfully")

	// Redirect to the widgets index page
	return c.Redirect(302, "/widgets")
}
