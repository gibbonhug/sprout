import { BoxProps } from '../interfaces';

/**
 * A box (Square)
 * @param props Props from JSON data. See BoxProps documentation
 * @returns Square div
 */
function Box(props: BoxProps) {
    return <div className='box'></div>;
}

export default Box;
