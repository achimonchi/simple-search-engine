package response

import "github.com/gofiber/fiber/v2"

func FiberResponse(ctx *fiber.Ctx, resp Response) error {
	return ctx.Status(resp.Status).JSON(resp)
}
