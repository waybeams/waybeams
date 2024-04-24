const std = @import("std");
const testing = std.testing;

const dom = struct {
    const Kind = enum {
        Text,
        Element,
    };

    const Text = struct {
        const kind = Kind.Text;
        const text: []const u8;
    };

    const Element = struct {
        const kind = Kind.Element;
        const tag: []const u8;
        const children: []const dom;
    };

    pub const create_text = (text: []const u8): dom => {
        return .Text{ .text = text };
    };

    pub const create_element = (tag: []const u8, children: []const dom): dom => {
        return .Element{ .tag = tag, .children = children };
    };
};

const render = (node: dom) void {
    switch (node) {
        case .Text(let text):
            std.debug.print(text);
        case .Element(let element):
            std.debug.print("<");
            std.debug.print(element.tag);
            std.debug.print(">");
            for (element.children) |child| {
                render(child);
            }
            std.debug.print("</");
            std.debug.print(element.tag);
            std.debug.print(">");
    }
};



test "basic add functionality" {
    try testing.expect(add(3, 7) == 10);

    const node = dom.create_element("div", &[
        dom.create_text("Hello, "),
        dom.create_element("strong", &[
            dom.create_text("world"),
        ]),
        dom.create_text("!"),
    ]);

    render(node);
}
