/**
 * @package ext/go-extension
 * Professional TitanPL extension with Golang native bridge.
 */

/** Main configuration interface */
export interface Config {}

/** Main Extension class */
export default class Extension {
  constructor(config?: Config);

  /**
   * Adds two numbers using the native Go implementation.
   */
  addNumber(n1: number, n2: number): Promise<number>;
}