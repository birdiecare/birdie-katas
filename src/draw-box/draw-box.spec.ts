import {drawBox} from "./draw-box";

describe('.drawBox', () => {
  test("should draw a 2x2 box", () => {
    expect(drawBox(2, 2)).toEqual(`
+--+
|  |
|  |
+--+`
    )
  });

  test("should draw a 6x3 box", () => {
    expect(drawBox(6, 3)).toEqual( `
+------+
|      |
|      |
|      |
+------+`)
  });
});

