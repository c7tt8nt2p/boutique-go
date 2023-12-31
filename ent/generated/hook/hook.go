// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/kx-boutique/ent/generated"
)

// The AuthFunc type is an adapter to allow the use of ordinary
// function as Auth mutator.
type AuthFunc func(context.Context, *generated.AuthMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f AuthFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.AuthMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.AuthMutation", m)
}

// The CartFunc type is an adapter to allow the use of ordinary
// function as Cart mutator.
type CartFunc func(context.Context, *generated.CartMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f CartFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.CartMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.CartMutation", m)
}

// The CartItemFunc type is an adapter to allow the use of ordinary
// function as CartItem mutator.
type CartItemFunc func(context.Context, *generated.CartItemMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f CartItemFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.CartItemMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.CartItemMutation", m)
}

// The CheckoutFunc type is an adapter to allow the use of ordinary
// function as Checkout mutator.
type CheckoutFunc func(context.Context, *generated.CheckoutMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f CheckoutFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.CheckoutMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.CheckoutMutation", m)
}

// The CheckoutItemFunc type is an adapter to allow the use of ordinary
// function as CheckoutItem mutator.
type CheckoutItemFunc func(context.Context, *generated.CheckoutItemMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f CheckoutItemFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.CheckoutItemMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.CheckoutItemMutation", m)
}

// The ProductFunc type is an adapter to allow the use of ordinary
// function as Product mutator.
type ProductFunc func(context.Context, *generated.ProductMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f ProductFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.ProductMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.ProductMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *generated.UserMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.UserMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, generated.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m generated.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m generated.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m generated.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op generated.Op) Condition {
	return func(_ context.Context, m generated.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m generated.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m generated.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m generated.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk generated.Hook, cond Condition) generated.Hook {
	return func(next generated.Mutator) generated.Mutator {
		return generated.MutateFunc(func(ctx context.Context, m generated.Mutation) (generated.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, generated.Delete|generated.Create)
func On(hk generated.Hook, op generated.Op) generated.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, generated.Update|generated.UpdateOne)
func Unless(hk generated.Hook, op generated.Op) generated.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) generated.Hook {
	return func(generated.Mutator) generated.Mutator {
		return generated.MutateFunc(func(context.Context, generated.Mutation) (generated.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []generated.Hook {
//		return []generated.Hook{
//			Reject(generated.Delete|generated.Update),
//		}
//	}
func Reject(op generated.Op) generated.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []generated.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...generated.Hook) Chain {
	return Chain{append([]generated.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() generated.Hook {
	return func(mutator generated.Mutator) generated.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...generated.Hook) Chain {
	newHooks := make([]generated.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
