package opts

import (
	"github.com/waybeams/waybeams/pkg/events"
	. "github.com/waybeams/waybeams/pkg/spec"
)

func BgColor(color uint) Option {
	return func(r ReadWriter) {
		r.SetBgColor(color)
	}
}

// ExcludeFromLayout will configure Spec.ExcludeFromLayout.
func ExcludeFromLayout(value bool) Option {
	return func(r ReadWriter) {
		r.SetExcludeFromLayout(value)
	}
}

// FlexHeight will set Spec.FlexHeight.
func FlexHeight(value float64) Option {
	return func(r ReadWriter) {
		r.SetFlexHeight(value)
	}
}

// FlexWidth will set Spec.FlexWidth.
func FlexWidth(value float64) Option {
	return func(r ReadWriter) {
		r.SetFlexWidth(value)
	}
}

func FontColor(color uint) Option {
	return func(r ReadWriter) {
		r.SetFontColor(color)
	}
}

func FontFace(face string) Option {
	return func(r ReadWriter) {
		r.SetFontFace(face)
	}
}

func FontSize(size float64) Option {
	return func(r ReadWriter) {
		r.SetFontSize(size)
	}
}

func Gutter(value float64) Option {
	return func(r ReadWriter) {
		r.SetGutter(value)
	}
}

func IsDisabled(isDisabled bool) Option {
	return func(r ReadWriter) {
		if isDisabled {
			r.SetState("disabled")
		}
	}
}

func IsFocusable(value bool) Option {
	return func(r ReadWriter) {
		r.SetIsFocusable(value)
	}
}

func IsMeasured(measured bool) Option {
	return func(r ReadWriter) {
		r.SetIsMeasured(measured)
	}
}

func IsText(value bool) Option {
	return func(r ReadWriter) {
		r.SetIsText(value)
	}
}

func IsTextInput(value bool) Option {
	return func(r ReadWriter) {
		r.SetIsTextInput(value)
	}
}

// HAlign will set Spec.HAlign.
func HAlign(align Alignment) Option {
	return func(r ReadWriter) {
		r.SetHAlign(align)
	}
}

// Height will set Spec.Height.
func Height(value float64) Option {
	return func(r ReadWriter) {
		r.SetHeight(value)
	}
}

func Key(value string) Option {
	return func(r ReadWriter) {
		r.SetKey(value)
	}
}

// LayoutType will set Spec.LayoutType.
func LayoutType(layoutType LayoutTypeValue) Option {
	return func(r ReadWriter) {
		r.SetLayoutType(layoutType)
	}
}

// MaxHeight will set Spec.MaxHeight.
func MaxHeight(value float64) Option {
	return func(r ReadWriter) {
		r.SetMaxHeight(value)
	}
}

// MaxWidth will set Spec.MaxWidth.
func MaxWidth(value float64) Option {
	return func(r ReadWriter) {
		r.SetMaxWidth(value)
	}
}

// MinHeight will set Spec.MinHeight.
func MinHeight(value float64) Option {
	return func(r ReadWriter) {
		r.SetMinHeight(value)
	}
}

// MinWidth will set Spec.MinWidth.
func MinWidth(value float64) Option {
	return func(r ReadWriter) {
		r.SetMinWidth(value)
	}
}

// Padding will set Spec.Padding, which will effectively set padding for
// all four sides as well (bottom, top, left, right, horizontal and vertical).
func Padding(value float64) Option {
	return func(r ReadWriter) {
		r.SetPadding(value)
	}
}

// PaddingBottom will set Spec.PaddingBottom.
func PaddingBottom(value float64) Option {
	return func(r ReadWriter) {
		r.SetPaddingBottom(value)
	}
}

// PaddingLeft will set Spec.PaddingLeft.
func PaddingLeft(value float64) Option {
	return func(r ReadWriter) {
		r.SetPaddingLeft(value)
	}
}

// PaddingRight will set Spec.PaddingRight.
func PaddingRight(value float64) Option {
	return func(r ReadWriter) {
		r.SetPaddingRight(value)
	}
}

// PaddingTop will set Spec.PaddingTop.
func PaddingTop(value float64) Option {
	return func(r ReadWriter) {
		r.SetPaddingTop(value)
	}
}

// PrefHeight will set Spec.PrefHeight.
func PrefHeight(value float64) Option {
	return func(r ReadWriter) {
		r.SetPrefHeight(value)
	}
}

// PrefWidth will set Spec.PrefWidth.
func PrefWidth(value float64) Option {
	return func(r ReadWriter) {
		r.SetPrefWidth(value)
	}
}

// Size will set Spec.Width and Spec.Height.
func Size(width, height float64) Option {
	return func(r ReadWriter) {
		r.SetWidth(width)
		r.SetHeight(height)
	}
}

func SpecName(name string) Option {
	return func(r ReadWriter) {
		r.SetSpecName(name)
	}
}

func StrokeColor(color uint) Option {
	return func(r ReadWriter) {
		r.SetStrokeColor(color)
	}
}

func StrokeSize(size float64) Option {
	return func(r ReadWriter) {
		r.SetStrokeSize(size)
	}
}

func Text(value string) Option {
	return func(r ReadWriter) {
		// TODO(lbayes): Sanitize text as user input values can be placed in here.
		// TODO(lbayes): Localize text using Localization map.
		r.SetText(value)
	}
}

// VAlign will set Spec.VAlign.
func VAlign(align Alignment) Option {
	return func(r ReadWriter) {
		r.SetVAlign(align)
	}
}

func View(view RenderHandler) Option {
	return func(r ReadWriter) {
		r.SetView(view)
	}
}

func Visible(visible bool) Option {
	return func(r ReadWriter) {
		r.SetVisible(visible)
	}
}

// Width will set Spec.Width.
func Width(value float64) Option {
	return func(r ReadWriter) {
		r.SetWidth(value)
	}
}

// X will set Spec.X.
func X(pos float64) Option {
	return func(r ReadWriter) {
		r.SetX(pos)
	}
}

// Y will set Spec.Y.
func Y(pos float64) Option {
	return func(r ReadWriter) {
		r.SetY(pos)
	}
}

//-------------------------------------------
// Special Adapters
//-------------------------------------------

func Empty() Option {
	return func(rw ReadWriter) {
		// noop
	}
}

// Child will add the provided ReadWriter as a child to the associated control.
func Child(child ReadWriter) Option {
	return func(rw ReadWriter) {
		rw.SetChildren(append(rw.Children(), child))
		child.SetParent(rw)
	}
}

// Children will call the provided factory function and append the returned children.
func Children(children []ReadWriter) Option {
	return func(rw ReadWriter) {
		rw.SetChildren(append(rw.Children(), children...))
		for _, child := range children {
			child.SetParent(rw)
		}
	}
}

func childIndexByKey(children []ReadWriter, key string) int {
	if key != "" {
		for index, child := range children {
			if child.Key() == key {
				return index
			}
		}
	}
	return -1
}

func Childf(factory func() ReadWriter) Option {
	return func(parent ReadWriter) {
		existingChildren := parent.Children()
		child := factory()
		// NOTE(lbayes): Likely opportunities to improve performance here!
		// We're traversing O(n2)
		matchedKeyIndex := childIndexByKey(existingChildren, child.Key())
		if matchedKeyIndex > -1 {
			existingChildren[matchedKeyIndex] = child
			parent.SetChildren(existingChildren)
		} else {
			parent.SetChildren(append(existingChildren, child))
		}

		child.SetParent(parent)
		child.SetFactory(factory)
	}
}

// Children will call the provided factory function and append the returned children.
func Childrenf(factory func() []ReadWriter) Option {
	return func(rw ReadWriter) {
		children := factory()
		rw.SetChildren(children)
		for _, child := range children {
			child.SetParent(rw)
			child.SetSiblingsFactory(factory)
		}
	}
}

// A Bag is a collection of Options that is itself an Option.
func Bag(opts ...Option) Option {
	return func(r ReadWriter) {
		for _, opt := range opts {
			opt(r)
		}
	}
}

// OptionsHandler will apply the provided options to the received Event target.
func OptionsHandler(options ...Option) events.EventHandler {
	return func(e events.Event) {
		target := e.Target().(ReadWriter)
		for _, option := range options {
			option(target)
		}
	}
}

//-------------------------------------------
// Event Helpers
//-------------------------------------------

// On will apply the provided handler to the provided event name.
func On(eventName string, handler events.EventHandler) Option {
	return func(r ReadWriter) {
		r.PushUnsub(r.On(eventName, handler))
	}
}

func OnClick(handler events.EventHandler) Option {
	return func(r ReadWriter) {
		r.PushUnsub(r.On(events.Clicked, handler))
	}
}

func OnEnterKey(handler events.EventHandler) Option {
	return func(r ReadWriter) {
		r.PushUnsub(r.On(events.EnterKeyReleased, handler))
	}
}

//-------------------------------------------
// State Helpers
//-------------------------------------------

// TODO(lbayes): Consider introducing AppendState and ReplaceState
func OnState(name string, options ...Option) Option {
	return func(r ReadWriter) {
		r.OnState(name, options...)
	}
}

func SetState(name string) Option {
	return func(r ReadWriter) {
		r.SetState(name)
	}
}
