const Generator = require("yeoman-generator");

module.exports = class extends Generator {
  constructor(args, opts) {
    // Calling the super constructor is important so our generator is correctly set up
    super(args, opts);

    this.argument("number", {
      type: String,
      required: true,
      desc: "Which number Project Euler challenge you are solving"
    });
    this.argument("function", {
      type: String,
      required: true,
      desc: "The name of the function you are creating"
    });

    this.log(`Challenge number: ${this.options.number}`);
    this.log(`Function name: ${this.options.function}`);
  }

  writing() {
    this.fs.copyTpl(
      this.templatePath("challenge"),
      this.destinationPath(`${this.options.number}/challenge.go`),
      { functionName: this.options.function }
    );
    this.fs.copyTpl(
      this.templatePath("challenge_test"),
      this.destinationPath(`${this.options.number}/challenge_test.go`),
      { functionName: this.options.function }
    );
  }

  end() {
    this.log("Done!");
  }
};
