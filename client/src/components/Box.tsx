import { useState } from 'react';
import { BoxProps } from '../interfaces';
import Flower from './Flower';
import Dirt from './Dirt'; 
import '../scss/react_components/Box.scss';

/**
 * A box (Square)
 * 
 * Boxes display Flower and Dirt child components
 * 
 * The user will have 8 Boxes to hold all of their flowers
 * 
 * @param props Props from JSON data. See BoxProps documentation
 * @returns Square div
 */
function Box(props: BoxProps) {
    const [isSelected, setIsSelected] = useState(false);
    const [selectedClassName, setSelectedClassName] =
        useState('box unselected-box');

    const thisFlower = props.flower;

    /**
     * We want to add a yellow border to boxes when they are clicked, but
     * only if they have a flower inside of them
     *
     * The border color is determined by the class of the box (see SCSS file)
     *
     * Clicking the box a second time will remove the border
     */
    const handleClick = () => {
        // Clicking an unselected box with a flower
        if (!isSelected && thisFlower) {
            setIsSelected(true);
            setSelectedClassName('box selected-box');
        } else if (isSelected) {
            // Clicking box a second time
            setIsSelected(false);
            setSelectedClassName('box unselected-box');
        }
    };

    return (
        <div className={selectedClassName} onClick={handleClick}>
            {thisFlower && (
                <Flower key={"flower" + thisFlower.id} {...thisFlower}></Flower>
            )}

            <Dirt key={"dirt" + props.id} id={props.id}></Dirt>
        </div>
    );
}

export default Box;
