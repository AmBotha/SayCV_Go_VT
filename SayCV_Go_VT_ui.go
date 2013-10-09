// This file was created by ui2walk and may be regenerated.
// DO NOT EDIT OR YOUR MODIFICATIONS WILL BE LOST!

package main

import (
	"github.com/lxn/walk"
)

type dialogUI struct {
	textEdit               *walk.TextEdit
	tabWidget              *walk.TabWidget
	displayTab             *walk.TabPage
	statusGroupBox         *walk.GroupBox
	toolButton_2           *walk.ToolButton
	toolButton_3           *walk.ToolButton
	toolButton_4           *walk.ToolButton
	toolButton_6           *walk.ToolButton
	toolButton_5           *walk.ToolButton
	toolButton_7           *walk.ToolButton
	toolButton_8           *walk.ToolButton
	toolButton_9           *walk.ToolButton
	toolButton_10          *walk.ToolButton
	stopBitsGroupBox_2     *walk.GroupBox
	checkBox_6             *walk.CheckBox
	baudComboBox_2         *walk.ComboBox
	baudLabel_2            *walk.Label
	toolButton_11          *walk.ToolButton
	toolButton_12          *walk.ToolButton
	bkgDisplayGroupBox     *walk.GroupBox
	displayAsGroupBox      *walk.GroupBox
	AsciiRadioButton       *walk.RadioButton
	AnsiRadioButton        *walk.RadioButton
	radioButton_4          *walk.RadioButton
	radioButton_3          *walk.RadioButton
	radioButton_6          *walk.RadioButton
	radioButton_5          *walk.RadioButton
	radioButton_8          *walk.RadioButton
	radioButton_7          *walk.RadioButton
	radioButton_9          *walk.RadioButton
	AsciiXRadioButton      *walk.RadioButton
	radioButton_12         *walk.RadioButton
	radioButton_14         *walk.RadioButton
	radioButton_13         *walk.RadioButton
	radioButton_11         *walk.RadioButton
	displayAsCheckBox      *walk.CheckBox
	rowsSpinBox            *walk.NumberEdit
	colsSpinBox            *walk.NumberEdit
	checkBox_5             *walk.CheckBox
	checkBox_7             *walk.CheckBox
	checkBox_2             *walk.CheckBox
	checkBox_4             *walk.CheckBox
	checkBox               *walk.CheckBox
	checkBox_3             *walk.CheckBox
	rowsLabel              *walk.Label
	colsLabel              *walk.Label
	portTab                *walk.TabPage
	portGroupBox           *walk.GroupBox
	baudLabel              *walk.Label
	baudComboBox           *walk.ComboBox
	portComboBox           *walk.ComboBox
	portLabel              *walk.Label
	pushButton             *walk.PushButton
	portCheckBox           *walk.CheckBox
	parityGroupBox         *walk.GroupBox
	parityNoneRadioButton  *walk.RadioButton
	parityOddRadioButton   *walk.RadioButton
	parityEvenRadioButton  *walk.RadioButton
	parityMarkRadioButton  *walk.RadioButton
	paritySpaceRadioButton *walk.RadioButton
	dataBitsGroupBox       *walk.GroupBox
	dataBits8RadioButton   *walk.RadioButton
	dataBits7RadioButton   *walk.RadioButton
	dataBits6RadioButton   *walk.RadioButton
	dataBits5RadioButton   *walk.RadioButton
	hwFlowCtrlGroupBox     *walk.GroupBox
	radioButton_27         *walk.RadioButton
	radioButton_29         *walk.RadioButton
	radioButton_28         *walk.RadioButton
	radioButton_30         *walk.RadioButton
	stopBitsGroupBox       *walk.GroupBox
	radioButton_26         *walk.RadioButton
	radioButton_25         *walk.RadioButton
	tab_3                  *walk.TabPage
	tab_4                  *walk.TabPage
	tab_5                  *walk.TabPage
	tab_6                  *walk.TabPage
	tab_7                  *walk.TabPage
	tab_8                  *walk.TabPage
	tab_9                  *walk.TabPage
	tab_10                 *walk.TabPage
}

func (w *Dialog) init(owner walk.Form) (err error) {
	if w.Dialog, err = walk.NewDialog(owner); err != nil {
		return err
	}

	succeeded := false
	defer func() {
		if !succeeded {
			w.Dispose()
		}
	}()

	var font *walk.Font
	if font == nil {
		font = nil
	}

	w.SetName("Dialog")
	if err := w.SetClientSize(walk.Size{800, 600}); err != nil {
		return err
	}
	if err := w.SetTitle(`Dialog`); err != nil {
		return err
	}

	// textEdit
	if w.ui.textEdit, err = walk.NewTextEdit(w); err != nil {
		return err
	}
	w.ui.textEdit.SetName("textEdit")
	if err := w.ui.textEdit.SetBounds(walk.Rectangle{0, 0, 800, 300}); err != nil {
		return err
	}

	// tabWidget
	if w.ui.tabWidget, err = walk.NewTabWidget(w); err != nil {
		return err
	}
	w.ui.tabWidget.SetName("tabWidget")
	if err := w.ui.tabWidget.SetBounds(walk.Rectangle{0, 300, 801, 271}); err != nil {
		return err
	}

	// displayTab
	if w.ui.displayTab, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.displayTab.SetName("displayTab")
	if err := w.ui.displayTab.SetTitle(`Display`); err != nil {
		return err
	}
	if err := w.ui.tabWidget.Pages().Add(w.ui.displayTab); err != nil {
		return err
	}

	// statusGroupBox
	if w.ui.statusGroupBox, err = walk.NewGroupBox(w.ui.displayTab); err != nil {
		return err
	}
	w.ui.statusGroupBox.SetName("statusGroupBox")
	if err := w.ui.statusGroupBox.SetBounds(walk.Rectangle{710, 0, 81, 241}); err != nil {
		return err
	}
	if err := w.ui.statusGroupBox.SetTitle(`Status`); err != nil {
		return err
	}

	// toolButton_2
	if w.ui.toolButton_2, err = walk.NewToolButton(w.ui.statusGroupBox); err != nil {
		return err
	}
	w.ui.toolButton_2.SetName("toolButton_2")
	if err := w.ui.toolButton_2.SetBounds(walk.Rectangle{10, 12, 59, 18}); err != nil {
		return err
	}
	if err := w.ui.toolButton_2.SetText(`Connected`); err != nil {
		return err
	}

	// toolButton_3
	if w.ui.toolButton_3, err = walk.NewToolButton(w.ui.statusGroupBox); err != nil {
		return err
	}
	w.ui.toolButton_3.SetName("toolButton_3")
	if err := w.ui.toolButton_3.SetBounds(walk.Rectangle{10, 38, 55, 18}); err != nil {
		return err
	}
	if err := w.ui.toolButton_3.SetText(`RXD(2)`); err != nil {
		return err
	}

	// toolButton_4
	if w.ui.toolButton_4, err = walk.NewToolButton(w.ui.statusGroupBox); err != nil {
		return err
	}
	w.ui.toolButton_4.SetName("toolButton_4")
	if err := w.ui.toolButton_4.SetBounds(walk.Rectangle{10, 64, 55, 18}); err != nil {
		return err
	}
	if err := w.ui.toolButton_4.SetText(`TXD(3)`); err != nil {
		return err
	}

	// toolButton_6
	if w.ui.toolButton_6, err = walk.NewToolButton(w.ui.statusGroupBox); err != nil {
		return err
	}
	w.ui.toolButton_6.SetName("toolButton_6")
	if err := w.ui.toolButton_6.SetBounds(walk.Rectangle{10, 90, 55, 18}); err != nil {
		return err
	}
	if err := w.ui.toolButton_6.SetText(`CTS(8)`); err != nil {
		return err
	}

	// toolButton_5
	if w.ui.toolButton_5, err = walk.NewToolButton(w.ui.statusGroupBox); err != nil {
		return err
	}
	w.ui.toolButton_5.SetName("toolButton_5")
	if err := w.ui.toolButton_5.SetBounds(walk.Rectangle{10, 116, 55, 18}); err != nil {
		return err
	}
	if err := w.ui.toolButton_5.SetText(`DCD(1)`); err != nil {
		return err
	}

	// toolButton_7
	if w.ui.toolButton_7, err = walk.NewToolButton(w.ui.statusGroupBox); err != nil {
		return err
	}
	w.ui.toolButton_7.SetName("toolButton_7")
	if err := w.ui.toolButton_7.SetBounds(walk.Rectangle{10, 142, 55, 18}); err != nil {
		return err
	}
	if err := w.ui.toolButton_7.SetText(`DSR(6)`); err != nil {
		return err
	}

	// toolButton_8
	if w.ui.toolButton_8, err = walk.NewToolButton(w.ui.statusGroupBox); err != nil {
		return err
	}
	w.ui.toolButton_8.SetName("toolButton_8")
	if err := w.ui.toolButton_8.SetBounds(walk.Rectangle{10, 168, 59, 18}); err != nil {
		return err
	}
	if err := w.ui.toolButton_8.SetText(`Ring(9)`); err != nil {
		return err
	}

	// toolButton_9
	if w.ui.toolButton_9, err = walk.NewToolButton(w.ui.statusGroupBox); err != nil {
		return err
	}
	w.ui.toolButton_9.SetName("toolButton_9")
	if err := w.ui.toolButton_9.SetBounds(walk.Rectangle{10, 194, 49, 18}); err != nil {
		return err
	}
	if err := w.ui.toolButton_9.SetText(`BREAK`); err != nil {
		return err
	}

	// toolButton_10
	if w.ui.toolButton_10, err = walk.NewToolButton(w.ui.statusGroupBox); err != nil {
		return err
	}
	w.ui.toolButton_10.SetName("toolButton_10")
	if err := w.ui.toolButton_10.SetBounds(walk.Rectangle{10, 220, 49, 18}); err != nil {
		return err
	}
	if err := w.ui.toolButton_10.SetText(`Error`); err != nil {
		return err
	}

	// stopBitsGroupBox_2
	if w.ui.stopBitsGroupBox_2, err = walk.NewGroupBox(w.ui.displayTab); err != nil {
		return err
	}
	w.ui.stopBitsGroupBox_2.SetName("stopBitsGroupBox_2")
	if err := w.ui.stopBitsGroupBox_2.SetBounds(walk.Rectangle{140, 110, 147, 68}); err != nil {
		return err
	}
	if err := w.ui.stopBitsGroupBox_2.SetTitle(`Data Frames`); err != nil {
		return err
	}

	// checkBox_6
	if w.ui.checkBox_6, err = walk.NewCheckBox(w.ui.stopBitsGroupBox_2); err != nil {
		return err
	}
	w.ui.checkBox_6.SetName("checkBox_6")
	if err := w.ui.checkBox_6.SetBounds(walk.Rectangle{10, 40, 54, 17}); err != nil {
		return err
	}
	if err := w.ui.checkBox_6.SetText(`Single`); err != nil {
		return err
	}

	// baudComboBox_2
	if w.ui.baudComboBox_2, err = walk.NewComboBox(w.ui.stopBitsGroupBox_2); err != nil {
		return err
	}
	w.ui.baudComboBox_2.SetName("baudComboBox_2")
	if err := w.ui.baudComboBox_2.SetBounds(walk.Rectangle{70, 10, 67, 20}); err != nil {
		return err
	}

	// baudLabel_2
	if w.ui.baudLabel_2, err = walk.NewLabel(w.ui.stopBitsGroupBox_2); err != nil {
		return err
	}
	w.ui.baudLabel_2.SetName("baudLabel_2")
	if err := w.ui.baudLabel_2.SetBounds(walk.Rectangle{10, 20, 30, 12}); err != nil {
		return err
	}
	if err := w.ui.baudLabel_2.SetText(`Bytes`); err != nil {
		return err
	}

	// toolButton_11
	if w.ui.toolButton_11, err = walk.NewToolButton(w.ui.stopBitsGroupBox_2); err != nil {
		return err
	}
	w.ui.toolButton_11.SetName("toolButton_11")
	if err := w.ui.toolButton_11.SetBounds(walk.Rectangle{90, 40, 43, 18}); err != nil {
		return err
	}
	if err := w.ui.toolButton_11.SetText(`Gulp`); err != nil {
		return err
	}

	// toolButton_12
	if w.ui.toolButton_12, err = walk.NewToolButton(w.ui.displayTab); err != nil {
		return err
	}
	w.ui.toolButton_12.SetName("toolButton_12")
	if err := w.ui.toolButton_12.SetBounds(walk.Rectangle{140, 220, 97, 18}); err != nil {
		return err
	}
	if err := w.ui.toolButton_12.SetText(`Terminal Font`); err != nil {
		return err
	}

	// bkgDisplayGroupBox
	if w.ui.bkgDisplayGroupBox, err = walk.NewGroupBox(w.ui.displayTab); err != nil {
		return err
	}
	w.ui.bkgDisplayGroupBox.SetName("bkgDisplayGroupBox")
	if err := w.ui.bkgDisplayGroupBox.SetBounds(walk.Rectangle{0, 0, 801, 251}); err != nil {
		return err
	}
	if err := w.ui.bkgDisplayGroupBox.SetTitle(``); err != nil {
		return err
	}

	// displayAsGroupBox
	if w.ui.displayAsGroupBox, err = walk.NewGroupBox(w.ui.bkgDisplayGroupBox); err != nil {
		return err
	}
	w.ui.displayAsGroupBox.SetName("displayAsGroupBox")
	if err := w.ui.displayAsGroupBox.SetBounds(walk.Rectangle{0, 0, 131, 251}); err != nil {
		return err
	}
	if err := w.ui.displayAsGroupBox.SetTitle(`Display As`); err != nil {
		return err
	}

	// AsciiRadioButton
	if w.ui.AsciiRadioButton, err = walk.NewRadioButton(w.ui.displayAsGroupBox); err != nil {
		return err
	}
	w.ui.AsciiRadioButton.SetName("AsciiRadioButton")
	if err := w.ui.AsciiRadioButton.SetBounds(walk.Rectangle{11, 23, 53, 16}); err != nil {
		return err
	}
	if err := w.ui.AsciiRadioButton.SetText(`Ascii`); err != nil {
		return err
	}

	// AnsiRadioButton
	if w.ui.AnsiRadioButton, err = walk.NewRadioButton(w.ui.displayAsGroupBox); err != nil {
		return err
	}
	w.ui.AnsiRadioButton.SetName("AnsiRadioButton")
	if err := w.ui.AnsiRadioButton.SetBounds(walk.Rectangle{11, 40, 47, 16}); err != nil {
		return err
	}
	if err := w.ui.AnsiRadioButton.SetText(`Ansi`); err != nil {
		return err
	}

	// radioButton_4
	if w.ui.radioButton_4, err = walk.NewRadioButton(w.ui.displayAsGroupBox); err != nil {
		return err
	}
	w.ui.radioButton_4.SetName("radioButton_4")
	if err := w.ui.radioButton_4.SetBounds(walk.Rectangle{11, 58, 83, 16}); err != nil {
		return err
	}
	if err := w.ui.radioButton_4.SetText(`Hex(Space)`); err != nil {
		return err
	}

	// radioButton_3
	if w.ui.radioButton_3, err = walk.NewRadioButton(w.ui.displayAsGroupBox); err != nil {
		return err
	}
	w.ui.radioButton_3.SetName("radioButton_3")
	if err := w.ui.radioButton_3.SetBounds(walk.Rectangle{11, 75, 89, 16}); err != nil {
		return err
	}
	if err := w.ui.radioButton_3.SetText(`Hex + Ascii`); err != nil {
		return err
	}

	// radioButton_6
	if w.ui.radioButton_6, err = walk.NewRadioButton(w.ui.displayAsGroupBox); err != nil {
		return err
	}
	w.ui.radioButton_6.SetName("radioButton_6")
	if err := w.ui.radioButton_6.SetBounds(walk.Rectangle{11, 92, 53, 16}); err != nil {
		return err
	}
	if err := w.ui.radioButton_6.SetText(`uint8`); err != nil {
		return err
	}

	// radioButton_5
	if w.ui.radioButton_5, err = walk.NewRadioButton(w.ui.displayAsGroupBox); err != nil {
		return err
	}
	w.ui.radioButton_5.SetName("radioButton_5")
	if err := w.ui.radioButton_5.SetBounds(walk.Rectangle{11, 110, 47, 16}); err != nil {
		return err
	}
	if err := w.ui.radioButton_5.SetText(`int8`); err != nil {
		return err
	}

	// radioButton_8
	if w.ui.radioButton_8, err = walk.NewRadioButton(w.ui.displayAsGroupBox); err != nil {
		return err
	}
	w.ui.radioButton_8.SetName("radioButton_8")
	if err := w.ui.radioButton_8.SetBounds(walk.Rectangle{11, 127, 41, 16}); err != nil {
		return err
	}
	if err := w.ui.radioButton_8.SetText(`Hex`); err != nil {
		return err
	}

	// radioButton_7
	if w.ui.radioButton_7, err = walk.NewRadioButton(w.ui.displayAsGroupBox); err != nil {
		return err
	}
	w.ui.radioButton_7.SetName("radioButton_7")
	if err := w.ui.radioButton_7.SetBounds(walk.Rectangle{11, 144, 53, 16}); err != nil {
		return err
	}
	if err := w.ui.radioButton_7.SetText(`int16`); err != nil {
		return err
	}

	// radioButton_9
	if w.ui.radioButton_9, err = walk.NewRadioButton(w.ui.displayAsGroupBox); err != nil {
		return err
	}
	w.ui.radioButton_9.SetName("radioButton_9")
	if err := w.ui.radioButton_9.SetBounds(walk.Rectangle{11, 162, 59, 16}); err != nil {
		return err
	}
	if err := w.ui.radioButton_9.SetText(`uint16`); err != nil {
		return err
	}

	// AsciiXRadioButton
	if w.ui.AsciiXRadioButton, err = walk.NewRadioButton(w.ui.displayAsGroupBox); err != nil {
		return err
	}
	w.ui.AsciiXRadioButton.SetName("AsciiXRadioButton")
	if err := w.ui.AsciiXRadioButton.SetBounds(walk.Rectangle{11, 179, 53, 16}); err != nil {
		return err
	}
	if err := w.ui.AsciiXRadioButton.SetText(`Ascii`); err != nil {
		return err
	}

	// radioButton_12
	if w.ui.radioButton_12, err = walk.NewRadioButton(w.ui.displayAsGroupBox); err != nil {
		return err
	}
	w.ui.radioButton_12.SetName("radioButton_12")
	if err := w.ui.radioButton_12.SetBounds(walk.Rectangle{11, 196, 59, 16}); err != nil {
		return err
	}
	if err := w.ui.radioButton_12.SetText(`Binary`); err != nil {
		return err
	}

	// radioButton_14
	if w.ui.radioButton_14, err = walk.NewRadioButton(w.ui.displayAsGroupBox); err != nil {
		return err
	}
	w.ui.radioButton_14.SetName("radioButton_14")
	if err := w.ui.radioButton_14.SetBounds(walk.Rectangle{11, 214, 59, 16}); err != nil {
		return err
	}
	if err := w.ui.radioButton_14.SetText(`Nibble`); err != nil {
		return err
	}

	// radioButton_13
	if w.ui.radioButton_13, err = walk.NewRadioButton(w.ui.displayAsGroupBox); err != nil {
		return err
	}
	w.ui.radioButton_13.SetName("radioButton_13")
	if err := w.ui.radioButton_13.SetBounds(walk.Rectangle{11, 231, 59, 16}); err != nil {
		return err
	}
	if err := w.ui.radioButton_13.SetText(`Flaot4`); err != nil {
		return err
	}

	// radioButton_11
	if w.ui.radioButton_11, err = walk.NewRadioButton(w.ui.displayAsGroupBox); err != nil {
		return err
	}
	w.ui.radioButton_11.SetName("radioButton_11")
	if err := w.ui.radioButton_11.SetBounds(walk.Rectangle{11, 248, 65, 16}); err != nil {
		return err
	}
	if err := w.ui.radioButton_11.SetText(`Hex CSV`); err != nil {
		return err
	}

	// displayAsCheckBox
	if w.ui.displayAsCheckBox, err = walk.NewCheckBox(w.ui.displayAsGroupBox); err != nil {
		return err
	}
	w.ui.displayAsCheckBox.SetName("displayAsCheckBox")
	if err := w.ui.displayAsCheckBox.SetBounds(walk.Rectangle{106, 23, 16, 16}); err != nil {
		return err
	}
	if err := w.ui.displayAsCheckBox.SetText(``); err != nil {
		return err
	}

	// rowsSpinBox
	if w.ui.rowsSpinBox, err = walk.NewNumberEdit(w.ui.bkgDisplayGroupBox); err != nil {
		return err
	}
	w.ui.rowsSpinBox.SetName("rowsSpinBox")
	if err := w.ui.rowsSpinBox.SetBounds(walk.Rectangle{250, 220, 44, 22}); err != nil {
		return err
	}

	// colsSpinBox
	if w.ui.colsSpinBox, err = walk.NewNumberEdit(w.ui.bkgDisplayGroupBox); err != nil {
		return err
	}
	w.ui.colsSpinBox.SetName("colsSpinBox")
	if err := w.ui.colsSpinBox.SetBounds(walk.Rectangle{300, 220, 44, 22}); err != nil {
		return err
	}

	// checkBox_5
	if w.ui.checkBox_5, err = walk.NewCheckBox(w.ui.bkgDisplayGroupBox); err != nil {
		return err
	}
	w.ui.checkBox_5.SetName("checkBox_5")
	w.ui.checkBox_5.SetEnabled(false)
	if err := w.ui.checkBox_5.SetBounds(walk.Rectangle{140, 70, 84, 17}); err != nil {
		return err
	}
	if err := w.ui.checkBox_5.SetText(`Big Endian`); err != nil {
		return err
	}

	// checkBox_7
	if w.ui.checkBox_7, err = walk.NewCheckBox(w.ui.bkgDisplayGroupBox); err != nil {
		return err
	}
	w.ui.checkBox_7.SetName("checkBox_7")
	if err := w.ui.checkBox_7.SetBounds(walk.Rectangle{360, 220, 84, 17}); err != nil {
		return err
	}
	if err := w.ui.checkBox_7.SetText(`Scrollback`); err != nil {
		return err
	}

	// checkBox_2
	if w.ui.checkBox_2, err = walk.NewCheckBox(w.ui.bkgDisplayGroupBox); err != nil {
		return err
	}
	w.ui.checkBox_2.SetName("checkBox_2")
	if err := w.ui.checkBox_2.SetBounds(walk.Rectangle{140, 30, 96, 17}); err != nil {
		return err
	}
	if err := w.ui.checkBox_2.SetText(`newLine mode`); err != nil {
		return err
	}

	// checkBox_4
	if w.ui.checkBox_4, err = walk.NewCheckBox(w.ui.bkgDisplayGroupBox); err != nil {
		return err
	}
	w.ui.checkBox_4.SetName("checkBox_4")
	if err := w.ui.checkBox_4.SetBounds(walk.Rectangle{200, 50, 54, 17}); err != nil {
		return err
	}
	if err := w.ui.checkBox_4.SetText(`7Bits`); err != nil {
		return err
	}

	// checkBox
	if w.ui.checkBox, err = walk.NewCheckBox(w.ui.bkgDisplayGroupBox); err != nil {
		return err
	}
	w.ui.checkBox.SetName("checkBox")
	if err := w.ui.checkBox.SetBounds(walk.Rectangle{140, 10, 90, 17}); err != nil {
		return err
	}
	if err := w.ui.checkBox.SetText(`Half Duplex`); err != nil {
		return err
	}

	// checkBox_3
	if w.ui.checkBox_3, err = walk.NewCheckBox(w.ui.bkgDisplayGroupBox); err != nil {
		return err
	}
	w.ui.checkBox_3.SetName("checkBox_3")
	if err := w.ui.checkBox_3.SetBounds(walk.Rectangle{140, 50, 60, 17}); err != nil {
		return err
	}
	if err := w.ui.checkBox_3.SetText(`Invert`); err != nil {
		return err
	}

	// rowsLabel
	if w.ui.rowsLabel, err = walk.NewLabel(w.ui.bkgDisplayGroupBox); err != nil {
		return err
	}
	w.ui.rowsLabel.SetName("rowsLabel")
	if err := w.ui.rowsLabel.SetBounds(walk.Rectangle{250, 200, 30, 12}); err != nil {
		return err
	}
	if err := w.ui.rowsLabel.SetText(`Rows`); err != nil {
		return err
	}

	// colsLabel
	if w.ui.colsLabel, err = walk.NewLabel(w.ui.bkgDisplayGroupBox); err != nil {
		return err
	}
	w.ui.colsLabel.SetName("colsLabel")
	if err := w.ui.colsLabel.SetBounds(walk.Rectangle{310, 200, 30, 12}); err != nil {
		return err
	}
	if err := w.ui.colsLabel.SetText(`Cols`); err != nil {
		return err
	}

	// portTab
	if w.ui.portTab, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.portTab.SetName("portTab")
	if err := w.ui.portTab.SetTitle(`Port`); err != nil {
		return err
	}
	if err := w.ui.tabWidget.Pages().Add(w.ui.portTab); err != nil {
		return err
	}

	// portGroupBox
	if w.ui.portGroupBox, err = walk.NewGroupBox(w.ui.portTab); err != nil {
		return err
	}
	w.ui.portGroupBox.SetName("portGroupBox")
	if err := w.ui.portGroupBox.SetBounds(walk.Rectangle{0, 0, 801, 251}); err != nil {
		return err
	}
	if err := w.ui.portGroupBox.SetTitle(``); err != nil {
		return err
	}

	// baudLabel
	if w.ui.baudLabel, err = walk.NewLabel(w.ui.portGroupBox); err != nil {
		return err
	}
	w.ui.baudLabel.SetName("baudLabel")
	if err := w.ui.baudLabel.SetBounds(walk.Rectangle{10, 30, 24, 12}); err != nil {
		return err
	}
	if err := w.ui.baudLabel.SetText(`Baud`); err != nil {
		return err
	}

	// baudComboBox
	if w.ui.baudComboBox, err = walk.NewComboBox(w.ui.portGroupBox); err != nil {
		return err
	}
	w.ui.baudComboBox.SetName("baudComboBox")
	if err := w.ui.baudComboBox.SetBounds(walk.Rectangle{40, 20, 74, 22}); err != nil {
		return err
	}

	// portComboBox
	if w.ui.portComboBox, err = walk.NewComboBox(w.ui.portGroupBox); err != nil {
		return err
	}
	w.ui.portComboBox.SetName("portComboBox")
	if err := w.ui.portComboBox.SetBounds(walk.Rectangle{160, 20, 74, 22}); err != nil {
		return err
	}

	// portLabel
	if w.ui.portLabel, err = walk.NewLabel(w.ui.portGroupBox); err != nil {
		return err
	}
	w.ui.portLabel.SetName("portLabel")
	if err := w.ui.portLabel.SetBounds(walk.Rectangle{130, 30, 24, 12}); err != nil {
		return err
	}
	if err := w.ui.portLabel.SetText(`Port`); err != nil {
		return err
	}

	// pushButton
	if w.ui.pushButton, err = walk.NewPushButton(w.ui.portGroupBox); err != nil {
		return err
	}
	w.ui.pushButton.SetName("pushButton")
	if err := w.ui.pushButton.SetBounds(walk.Rectangle{250, 20, 75, 23}); err != nil {
		return err
	}
	if err := w.ui.pushButton.SetText(`Open`); err != nil {
		return err
	}

	// portCheckBox
	if w.ui.portCheckBox, err = walk.NewCheckBox(w.ui.portGroupBox); err != nil {
		return err
	}
	w.ui.portCheckBox.SetName("portCheckBox")
	if err := w.ui.portCheckBox.SetBounds(walk.Rectangle{340, 20, 13, 13}); err != nil {
		return err
	}
	if err := w.ui.portCheckBox.SetText(``); err != nil {
		return err
	}

	// parityGroupBox
	if w.ui.parityGroupBox, err = walk.NewGroupBox(w.ui.portGroupBox); err != nil {
		return err
	}
	w.ui.parityGroupBox.SetName("parityGroupBox")
	if err := w.ui.parityGroupBox.SetBounds(walk.Rectangle{10, 70, 73, 143}); err != nil {
		return err
	}
	if err := w.ui.parityGroupBox.SetTitle(`Parity`); err != nil {
		return err
	}

	// parityNoneRadioButton
	if w.ui.parityNoneRadioButton, err = walk.NewRadioButton(w.ui.parityGroupBox); err != nil {
		return err
	}
	w.ui.parityNoneRadioButton.SetName("parityNoneRadioButton")
	if err := w.ui.parityNoneRadioButton.SetBounds(walk.Rectangle{11, 23, 47, 17}); err != nil {
		return err
	}
	if err := w.ui.parityNoneRadioButton.SetText(`None`); err != nil {
		return err
	}

	// parityOddRadioButton
	if w.ui.parityOddRadioButton, err = walk.NewRadioButton(w.ui.parityGroupBox); err != nil {
		return err
	}
	w.ui.parityOddRadioButton.SetName("parityOddRadioButton")
	if err := w.ui.parityOddRadioButton.SetBounds(walk.Rectangle{11, 46, 41, 17}); err != nil {
		return err
	}
	if err := w.ui.parityOddRadioButton.SetText(`Odd`); err != nil {
		return err
	}

	// parityEvenRadioButton
	if w.ui.parityEvenRadioButton, err = walk.NewRadioButton(w.ui.parityGroupBox); err != nil {
		return err
	}
	w.ui.parityEvenRadioButton.SetName("parityEvenRadioButton")
	if err := w.ui.parityEvenRadioButton.SetBounds(walk.Rectangle{11, 69, 47, 17}); err != nil {
		return err
	}
	if err := w.ui.parityEvenRadioButton.SetText(`Even`); err != nil {
		return err
	}

	// parityMarkRadioButton
	if w.ui.parityMarkRadioButton, err = walk.NewRadioButton(w.ui.parityGroupBox); err != nil {
		return err
	}
	w.ui.parityMarkRadioButton.SetName("parityMarkRadioButton")
	if err := w.ui.parityMarkRadioButton.SetBounds(walk.Rectangle{11, 92, 47, 17}); err != nil {
		return err
	}
	if err := w.ui.parityMarkRadioButton.SetText(`Mark`); err != nil {
		return err
	}

	// paritySpaceRadioButton
	if w.ui.paritySpaceRadioButton, err = walk.NewRadioButton(w.ui.parityGroupBox); err != nil {
		return err
	}
	w.ui.paritySpaceRadioButton.SetName("paritySpaceRadioButton")
	if err := w.ui.paritySpaceRadioButton.SetBounds(walk.Rectangle{11, 115, 53, 17}); err != nil {
		return err
	}
	if err := w.ui.paritySpaceRadioButton.SetText(`Space`); err != nil {
		return err
	}

	// dataBitsGroupBox
	if w.ui.dataBitsGroupBox, err = walk.NewGroupBox(w.ui.portGroupBox); err != nil {
		return err
	}
	w.ui.dataBitsGroupBox.SetName("dataBitsGroupBox")
	if err := w.ui.dataBitsGroupBox.SetBounds(walk.Rectangle{87, 70, 79, 143}); err != nil {
		return err
	}
	if err := w.ui.dataBitsGroupBox.SetTitle(`Data Bits`); err != nil {
		return err
	}

	// dataBits8RadioButton
	if w.ui.dataBits8RadioButton, err = walk.NewRadioButton(w.ui.dataBitsGroupBox); err != nil {
		return err
	}
	w.ui.dataBits8RadioButton.SetName("dataBits8RadioButton")
	if err := w.ui.dataBits8RadioButton.SetBounds(walk.Rectangle{11, 27, 55, 17}); err != nil {
		return err
	}
	if err := w.ui.dataBits8RadioButton.SetText(`8 bits`); err != nil {
		return err
	}

	// dataBits7RadioButton
	if w.ui.dataBits7RadioButton, err = walk.NewRadioButton(w.ui.dataBitsGroupBox); err != nil {
		return err
	}
	w.ui.dataBits7RadioButton.SetName("dataBits7RadioButton")
	if err := w.ui.dataBits7RadioButton.SetBounds(walk.Rectangle{11, 54, 55, 17}); err != nil {
		return err
	}
	if err := w.ui.dataBits7RadioButton.SetText(`7 bits`); err != nil {
		return err
	}

	// dataBits6RadioButton
	if w.ui.dataBits6RadioButton, err = walk.NewRadioButton(w.ui.dataBitsGroupBox); err != nil {
		return err
	}
	w.ui.dataBits6RadioButton.SetName("dataBits6RadioButton")
	if err := w.ui.dataBits6RadioButton.SetBounds(walk.Rectangle{11, 81, 55, 17}); err != nil {
		return err
	}
	if err := w.ui.dataBits6RadioButton.SetText(`6 bits`); err != nil {
		return err
	}

	// dataBits5RadioButton
	if w.ui.dataBits5RadioButton, err = walk.NewRadioButton(w.ui.dataBitsGroupBox); err != nil {
		return err
	}
	w.ui.dataBits5RadioButton.SetName("dataBits5RadioButton")
	if err := w.ui.dataBits5RadioButton.SetBounds(walk.Rectangle{11, 108, 55, 17}); err != nil {
		return err
	}
	if err := w.ui.dataBits5RadioButton.SetText(`5 bits`); err != nil {
		return err
	}

	// hwFlowCtrlGroupBox
	if w.ui.hwFlowCtrlGroupBox, err = walk.NewGroupBox(w.ui.portGroupBox); err != nil {
		return err
	}
	w.ui.hwFlowCtrlGroupBox.SetName("hwFlowCtrlGroupBox")
	if err := w.ui.hwFlowCtrlGroupBox.SetBounds(walk.Rectangle{170, 126, 170, 80}); err != nil {
		return err
	}
	if err := w.ui.hwFlowCtrlGroupBox.SetTitle(`Hardware Flow Control`); err != nil {
		return err
	}

	// radioButton_27
	if w.ui.radioButton_27, err = walk.NewRadioButton(w.ui.hwFlowCtrlGroupBox); err != nil {
		return err
	}
	w.ui.radioButton_27.SetName("radioButton_27")
	if err := w.ui.radioButton_27.SetBounds(walk.Rectangle{11, 23, 47, 17}); err != nil {
		return err
	}
	if err := w.ui.radioButton_27.SetText(`None`); err != nil {
		return err
	}

	// radioButton_29
	if w.ui.radioButton_29, err = walk.NewRadioButton(w.ui.hwFlowCtrlGroupBox); err != nil {
		return err
	}
	w.ui.radioButton_29.SetName("radioButton_29")
	if err := w.ui.radioButton_29.SetBounds(walk.Rectangle{11, 52, 65, 17}); err != nil {
		return err
	}
	if err := w.ui.radioButton_29.SetText(`DTR/DSR`); err != nil {
		return err
	}

	// radioButton_28
	if w.ui.radioButton_28, err = walk.NewRadioButton(w.ui.hwFlowCtrlGroupBox); err != nil {
		return err
	}
	w.ui.radioButton_28.SetName("radioButton_28")
	if err := w.ui.radioButton_28.SetBounds(walk.Rectangle{82, 23, 65, 17}); err != nil {
		return err
	}
	if err := w.ui.radioButton_28.SetText(`RTS/CTS`); err != nil {
		return err
	}

	// radioButton_30
	if w.ui.radioButton_30, err = walk.NewRadioButton(w.ui.hwFlowCtrlGroupBox); err != nil {
		return err
	}
	w.ui.radioButton_30.SetName("radioButton_30")
	if err := w.ui.radioButton_30.SetBounds(walk.Rectangle{82, 52, 77, 17}); err != nil {
		return err
	}
	if err := w.ui.radioButton_30.SetText(`RS485-rts`); err != nil {
		return err
	}

	// stopBitsGroupBox
	if w.ui.stopBitsGroupBox, err = walk.NewGroupBox(w.ui.portGroupBox); err != nil {
		return err
	}
	w.ui.stopBitsGroupBox.SetName("stopBitsGroupBox")
	if err := w.ui.stopBitsGroupBox.SetBounds(walk.Rectangle{170, 70, 146, 51}); err != nil {
		return err
	}
	if err := w.ui.stopBitsGroupBox.SetTitle(`Stop Bits`); err != nil {
		return err
	}

	// radioButton_26
	if w.ui.radioButton_26, err = walk.NewRadioButton(w.ui.stopBitsGroupBox); err != nil {
		return err
	}
	w.ui.radioButton_26.SetName("radioButton_26")
	if err := w.ui.radioButton_26.SetBounds(walk.Rectangle{11, 23, 59, 17}); err != nil {
		return err
	}
	if err := w.ui.radioButton_26.SetText(`1 bits`); err != nil {
		return err
	}

	// radioButton_25
	if w.ui.radioButton_25, err = walk.NewRadioButton(w.ui.stopBitsGroupBox); err != nil {
		return err
	}
	w.ui.radioButton_25.SetName("radioButton_25")
	if err := w.ui.radioButton_25.SetBounds(walk.Rectangle{76, 23, 59, 17}); err != nil {
		return err
	}
	if err := w.ui.radioButton_25.SetText(`2 bits`); err != nil {
		return err
	}

	// tab_3
	if w.ui.tab_3, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tab_3.SetName("tab_3")
	if err := w.ui.tab_3.SetTitle(`Capture`); err != nil {
		return err
	}
	if err := w.ui.tabWidget.Pages().Add(w.ui.tab_3); err != nil {
		return err
	}

	// tab_4
	if w.ui.tab_4, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tab_4.SetName("tab_4")
	if err := w.ui.tab_4.SetTitle(`Pins`); err != nil {
		return err
	}
	if err := w.ui.tabWidget.Pages().Add(w.ui.tab_4); err != nil {
		return err
	}

	// tab_5
	if w.ui.tab_5, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tab_5.SetName("tab_5")
	if err := w.ui.tab_5.SetTitle(`Send`); err != nil {
		return err
	}
	if err := w.ui.tabWidget.Pages().Add(w.ui.tab_5); err != nil {
		return err
	}

	// tab_6
	if w.ui.tab_6, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tab_6.SetName("tab_6")
	if err := w.ui.tab_6.SetTitle(`Echo Port`); err != nil {
		return err
	}
	if err := w.ui.tabWidget.Pages().Add(w.ui.tab_6); err != nil {
		return err
	}

	// tab_7
	if w.ui.tab_7, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tab_7.SetName("tab_7")
	if err := w.ui.tab_7.SetTitle(`I2C`); err != nil {
		return err
	}
	if err := w.ui.tabWidget.Pages().Add(w.ui.tab_7); err != nil {
		return err
	}

	// tab_8
	if w.ui.tab_8, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tab_8.SetName("tab_8")
	if err := w.ui.tab_8.SetTitle(`I2C-2`); err != nil {
		return err
	}
	if err := w.ui.tabWidget.Pages().Add(w.ui.tab_8); err != nil {
		return err
	}

	// tab_9
	if w.ui.tab_9, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tab_9.SetName("tab_9")
	if err := w.ui.tab_9.SetTitle(`I2CMisc`); err != nil {
		return err
	}
	if err := w.ui.tabWidget.Pages().Add(w.ui.tab_9); err != nil {
		return err
	}

	// tab_10
	if w.ui.tab_10, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tab_10.SetName("tab_10")
	if err := w.ui.tab_10.SetTitle(`Misc`); err != nil {
		return err
	}
	if err := w.ui.tabWidget.Pages().Add(w.ui.tab_10); err != nil {
		return err
	}

	// Tab order

	succeeded = true

	return nil
}
