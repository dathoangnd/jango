<template>
  <div class="jg-menu jg-btn-back">
    <div class="jg-icon"></div>
  </div>
</template>

<style lang="scss" scoped>
  .jg-btn-back{
    width: 24px;
    height: 24px;
    transition-duration: 0.5s;
    
    .jg-icon {
      transition-duration: 0.5s;
      position: absolute;
      height: 2px;
      width: 24px;
      top: 24px;
      background-color: blue;
      
      &:before{
        transition-duration: 0.5s;
        position: absolute;
        width: 24px;
        height: 2px;
        background-color: blue;
        content: "";
        top: -5px;
      }
      
      &:after{
        transition-duration: 0.5s;
        position: absolute;
        width: 24px;
        height: 2px;
        background-color: blue;
        content: "";
        top: 5px;
      }
    }
    
    &.jg-open {
      .jg-icon {
        transition: 0.5s;
        background-color: red;
        
        &:before{
          transform: rotateZ(45deg) scaleX(0.75) translate(10px, -3px);
          background-color: red;
        }

        &:after{
          transform: rotateZ(-45deg) scaleX(0.75) translate(10px, 3px);
          background-color: red;
        }
      }
    }
    &:hover {
      cursor: pointer;
    }
  }
</style>
