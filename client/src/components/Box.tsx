import { BoxProps } from '../interfaces';
import '../scss/react_components/Box.scss';

/**
 * A box (Square)
 * @param props Props from JSON data. See BoxProps documentation
 * @returns Square div
 */
function Box(props: BoxProps) {
    return <div className="box">Box {props.id}</div>;
}

export default Box;
