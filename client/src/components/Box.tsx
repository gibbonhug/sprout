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
 * @returns Square div which holds a Flower (depending on data) and a plot of Dirt
 */
function Box(props: any) {
    // Grab state passed from App component
    const idSelectedFlowerBox = props.idSelectedFlowerBox;
    const setIDSelectedFlowerBox = props.setIDSelectedFlowerBox;

    // Class name state used to determine styling (see Box.scss)
    const [selectedClassName, setSelectedClassName] =
        useState('box unselected-box');

    const thisFlower = props.flower;

    /**
     * Func to determine if this Box is selected
     * @param idSelectedFlowerBox The idSelectedFlowerBox prop passed from the App component
     * @return Bool of whether this Box is selected (matches ID of selected flower state)
     */
    const isThisBoxSelected = (idSelectedFlowerBox: any) => {
        return props.id === idSelectedFlowerBox;
    }

    /**
     * We want to add a yellow border to boxes when they are clicked, but
     * only if they have a flower inside of them
     *
     * The border color is determined by the class of the box (see SCSS file)
     *
     * Clicking the box a second time will remove the border
     */
    const handleClick = () => {
        if (!isThisBoxSelected(idSelectedFlowerBox) && thisFlower) {
            // Clicking an unselected box with a flower: Select this box
            setIDSelectedFlowerBox(props.id)
            setSelectedClassName('box selected-box');
        } else if (isThisBoxSelected(idSelectedFlowerBox)) {
            // Clicking selected box a second time: Unselect
            setIDSelectedFlowerBox(null);
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
