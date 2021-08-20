    /**
     * @param {string} digits
     * @return {string[]}
     */
    var letterCombinations = function (digits) {
        debugger;
        if (digits.length === 0) return [];

        let output_arr = new LinkedList();
        output_arr.add("");

        let char_map = ["0", "1", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"];

        for (let i = 0; i < digits.length; i++) {
            let index = parseInt(digits[i]);

            while (output_arr.peek().length === i) {
                console.log(`output_arr.peek(): ${output_arr.peek()}`);
                // console.log(output_arr.toArray());

                let permutation = output_arr.remove();
                console.log(`permutation: ${permutation}`);

                for (c of char_map[index].split('')) {
                    console.log(`c: ${c}`);
                    console.log(`permutation + c: ${permutation + c}`);
                    output_arr.add(permutation + c);
                }
                console.log(output_arr.toArray());
            }
        }

        return output_arr.toArray();
    };

    class Node {
        constructor(data, next = null) {
            this.data = data;
            this.next = next;
        }
    }

    class LinkedList {
        constructor() {
            this.head = null;
            this.size = 0;
        }

        add(data) {
            let current = this.head;

            if (current === null) {
                console.log(`added ${data}`);
                this.head = new Node(data);
            } else {
                while (current.next !== null) {
                    current = current.next;
                }

                console.log(`added ${data}`);
                let n = new Node(data);
                current.next = n;
            }

            this.size++;
        }

        peek() {
            return this.head.data;
        }

        remove() {
            let current = this.head;

            if (!this.head) {
                return;
            }

            this.head = this.head.next;

            this.size--;
            return current.data;
        }

        toArray() {
            let arr = [];

            let current = this.head;
            while (current) {
                console.log(`current.data: ${current.data}`);
                arr.push(current.data);
                current = current.next;
            }

            return arr;
        }
    }