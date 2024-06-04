pragma solidity ^0.5.16;
pragma experimental ABIEncoderV2;

contract write{
    uint verify_pi_fs_res=0;

    function test(bytes memory x1,bytes memory x2,uint rM,bytes memory y1)public{
        return;
        bytes memory input1=abi.encodePacked(x1,rM);
        bytes32 output;
        uint len_in=input1.length;
        uint len_out=output.length;
        assembly{
            if iszero(call(not(0),0x07,0,add(input1,0x20),len_in,output,len_out)){
                // revert(0,0)
            }
        }
        bytes memory input2=abi.encodePacked(output,rM,x2,rM);
        len_in=input2.length;
        uint[1] memory res;


        assembly{
            if iszero(call(not(0),0x08,0,add(input2,0x20),len_in,res,0x20)){
                // revert(0,0)
            }
        }
    }
}