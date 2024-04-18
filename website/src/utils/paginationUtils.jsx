import _ from "lodash";

export const returnPaginationRange = (totalPage, page, support) => {


    let totalPageNoInArray = 5 + support;
    if(totalPageNoInArray >= totalPage){
        return _.range(1, totalPage + 1);
    }
    let leftSiblingIndex = Math.max(page - support ,1);
    let rightSiblingIndex = Math.min(page + support , totalPage);

    let showLeftDots = leftSiblingIndex > 2;
    let showRightDots = rightSiblingIndex < totalPage -2;

    if(!showLeftDots && showRightDots){
        let leftItemsCount = 3 + 2 * support;
        let leftRange = _.range(1, leftItemsCount + 1);
        return [...leftRange, " ...", totalPage];
    }
    else if(showLeftDots && !showRightDots){
        let rightItemsCount = 3 + 2 * support;
        let rightRange = _.range(totalPage - rightItemsCount + 1, totalPage + 1);
        return [1, "... ", ...rightRange];
    }
    else{
        let middleRange = _.range(leftSiblingIndex, rightSiblingIndex + 1);
        return [1, "... ", ...middleRange, " ...", totalPage];
    }
}