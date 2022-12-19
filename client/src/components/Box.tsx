import { BoxProps } from '../interfaces';
import Flower from './Flower';
import '../scss/react_components/Box.scss';

/**
 * A box (Square)
 * @param props Props from JSON data. See BoxProps documentation
 * @returns Square div
 */
function Box(props: BoxProps) {
    const thisFlower = props.flower;

    return (
        <div className="box">
            Box {props.id}
            {thisFlower && <Flower key={thisFlower.id} {...thisFlower}></Flower>}
        </div>
    );
}

export default Box;
