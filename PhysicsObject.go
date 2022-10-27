package main

type PhysicsObject struct {
	pos, lastPos                  *Vec2
	vel, lastVel                  *Vec2
	collider                      *Rect
	colliderOffset                *Vec2
	isGrounded, wasGrounded       bool
	isAtCeiling, wasAtCeiling     bool
	isOnRightWall, wasOnRightWall bool
	isOnLeftWall, wasOnLeftWall   bool
}

func newPhysicsObject(pos Vec2, colliderWidth, colliderHeight float64, colliderOffset Vec2) *PhysicsObject {
	po := &PhysicsObject{}
	po.pos.x = pos.x
	po.pos.y = pos.y
	po.colliderOffset.x = colliderOffset.x
	po.colliderOffset.y = colliderOffset.y
	po.collider.w = colliderWidth
	po.collider.h = colliderHeight

	return po
}

func (po *PhysicsObject) UpdatePhysics(dt float64) {
	po.lastPos.x = po.pos.x
	po.lastPos.y = po.pos.y
	po.lastVel.x = po.vel.x
	po.lastVel.y = po.vel.y
	po.wasGrounded = po.isGrounded
	po.wasAtCeiling = po.isAtCeiling
	po.wasOnRightWall = po.isOnRightWall
	po.wasOnLeftWall = po.isOnLeftWall

	po.pos.x += po.vel.x * dt
	po.pos.y += po.vel.y * dt

	if po.pos.y > 415 {
		po.pos.y = 415
		po.isGrounded = true
	} else {
		po.isGrounded = false
	}

	po.collider.x = po.pos.x + po.colliderOffset.x
	po.collider.y = po.pos.y + po.colliderOffset.y
}
