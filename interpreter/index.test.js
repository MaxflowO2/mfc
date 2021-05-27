const Interpreter = require('./index');
const {
  STOP,
  ADD,
  SUB,
  MUL,
  DIV,
  PUSH,
  LT,
  GT,
  EQ,
  AND,
  OR,
  JUMP,
  JUMPI
} = Interpreter.OPCODE_MAP;

describe('Interpreter', () => {
  describe('runCode()', () => {
    describe('and the code includes ADD', () => {
      it('adds two values', () => {
        expect (
          new Interpreter().runCode([PUSH, 6, PUSH, 3, ADD, STOP])
        ).toEqual(9);
      });
    });
    describe('and the code includes SUB', () => {
      it('subtracts two values', () => {
        expect (
          new Interpreter().runCode([PUSH, 6, PUSH, 3, SUB, STOP])
        ).toEqual(-3);
      });
    });
    describe('and the code includes MUL', () => {
      it('multiplies two values', () => {
        expect (
          new Interpreter().runCode([PUSH, 6, PUSH, 3, MUL, STOP])
        ).toEqual(18);
      });
    });
    describe('and the code includes DIV', () => {
      it('divides two values', () => {
        expect (
          new Interpreter().runCode([PUSH, 6, PUSH, 3, DIV, STOP])
        ).toEqual(0.5);
      });
    });
  });
});
