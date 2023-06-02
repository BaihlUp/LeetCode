//给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并
//返回其根节点。 
//
// 
//
// 示例 1: 
// 
// 
//输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//输出: [3,9,20,null,null,15,7]
// 
//
// 示例 2: 
//
// 
//输入: preorder = [-1], inorder = [-1]
//输出: [-1]
// 
//
// 
//
// 提示: 
//
// 
// 1 <= preorder.length <= 3000 
// inorder.length == preorder.length 
// -3000 <= preorder[i], inorder[i] <= 3000 
// preorder 和 inorder 均 无重复 元素 
// inorder 均出现在 preorder 
// preorder 保证 为二叉树的前序遍历序列 
// inorder 保证 为二叉树的中序遍历序列 
// 
//
// Related Topics 树 数组 哈希表 分治 二叉树 👍 1978 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
        return help(preorder, 0, preorder.size()-1, inorder, 0, inorder.size()-1);
    }
    TreeNode* help(vector<int>& preorder, int prestart, int preend, vector<int> &inorder, int instart, int inend) {
        if(prestart > preend) return NULL;

        int rootVal = preorder[prestart];
        int index = 0;
        for(int i = instart; i <= inend; i++) {
            if(rootVal == inorder[i])
                index = i;
        }
        int leftSize = index-instart;
        TreeNode* root = new TreeNode(rootVal);
        root->left = help(preorder, prestart+1, prestart+leftSize, inorder, instart, index-1);
        root->right = help(preorder, prestart+leftSize+1, preend, inorder, index+1, inend);
        return root;
    }
};
//leetcode submit region end(Prohibit modification and deletion)
